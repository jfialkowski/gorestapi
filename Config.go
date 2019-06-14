package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

//Config is the main struct for the JSON objects in config
type Config struct {
	Name      string
	DbConfig  dbConfig
	TLSConfig TLSConfig
}

//dbConfig is the Database configuration related JSON Objects
type dbConfig struct {
	Host         string
	Databasename string
	Port         string
	Username     string
	Password     string
}

//TLSConfig is the TLS configuration related JSON Objects
type TLSConfig struct {
	Certificate string
	Key         string
	Passphrase  string
	ServerPort  string
}

//DBhost variable
var DBhost string

//DBport variable
var DBport string

//DBuser variable
var DBuser string

//DBpass variable
var DBpass string

//DBname variable
var DBname string

//TLSCert is the wbserver certificate
var TLSCert []byte

//TLSKey is the wbserver private key
var TLSKey []byte

//TLSPass is the password to key if found, nil otherwise.
var TLSPass string

//ServerPort is the port you want the app to run on
var ServerPort string

//Gets AWS Credentials from ENV
func getKeys() (string, string) {

	keyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	return keyID, secretAccessKey
}

//Gets the config JSON from AWS secretsmanager and returns it as a string
func getConfig() []byte {
	keyID, secretAccessKey := getKeys()
	secretName := "gorestapiconfig"
	config := ""
	//region := "us-east-1"

	//Create a Secrets Manager client

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(keyID, secretAccessKey, ""),
	})

	if err != nil {
		fmt.Println(err)
	}

	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
		config = secretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			os.Exit(42)
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
		config = decodedBinarySecret

	}
	log.Println("Retrieved Configuration From Secrets Manager")

	return []byte(config)

}

//LoadConfig parses the JSON config fetched from getConfig and sets the variables the app needs to function
func LoadConfig() {
	config := getConfig()
	if json.Valid(config) {
		log.Println("Configuration Loaded")
	} else {
		panic("Could Not load Config because of invalid JSON")
	}

	var jsonVaules Config
	json.Unmarshal(config, &jsonVaules)
	DBhost = jsonVaules.DbConfig.Host
	DBport = jsonVaules.DbConfig.Port
	DBuser = jsonVaules.DbConfig.Username
	DBpass = jsonVaules.DbConfig.Password
	DBname = jsonVaules.DbConfig.Databasename
	TLSCert = []byte(jsonVaules.TLSConfig.Certificate)
	TLSKey = []byte(jsonVaules.TLSConfig.Key)
	TLSPass = jsonVaules.TLSConfig.Passphrase
	ServerPort = jsonVaules.TLSConfig.ServerPort

}
