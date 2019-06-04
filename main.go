package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func find(obj interface{}, key string) (interface{}, bool) {
	//if the argument is not a map, ignore it
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}

	for k, v := range mobj {
		//key match, return value
		if k == key {
			return v, true
		}

		//if the value is a map, search recursively
		if m, ok := v.(map[string]interface{}); ok {
			if res, ok := find(m, key); ok {
				return res, true
			}
		}
		//if the value is an array, search recursively
		//from each element
		if va, ok := v.([]interface{}); ok {
			for _, a := range va {
				if res, ok := find(a, key); ok {
					return res, true
				}
			}
		}
	}

	//element not found
	return nil, false
}

func main() {
	dbHost := ""
	dbPort := ""
	dbUser := ""
	dbPass := ""
	dbName := ""
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	router := NewRouter()
	config := getConfig()
	configMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(config), &configMap)
	if err != nil {
		panic(err)
	}
	if host, ok := find(configMap, "host"); ok {
		switch v := host.(type) {
		case string:
			dbHost = v
		case fmt.Stringer:

		case int:

		default:

		}
	}
	if port, ok := find(configMap, "port"); ok {
		switch v := port.(type) {
		case string:
			dbPort = v
		case fmt.Stringer:

		case int:
			dbPort = string(v)
		default:

		}
	}
	if username, ok := find(configMap, "username"); ok {
		switch v := username.(type) {
		case string:
			dbUser = v
		case fmt.Stringer:

		case int:

		default:

		}
	}
	if password, ok := find(configMap, "password"); ok {
		switch v := password.(type) {
		case string:
			dbPass = v
		case fmt.Stringer:

		case int:

		default:

		}
	}
	if databasename, ok := find(configMap, "databasename"); ok {
		switch v := databasename.(type) {
		case string:
			dbName = v
		case fmt.Stringer:

		case int:

		default:

		}
	}

	fmt.Println("Host: " + dbHost)
	fmt.Println("Port: " + dbPort)
	fmt.Println("User: " + dbUser)
	fmt.Println("Pass: " + dbPass)
	fmt.Println("Name: " + dbName)
	log.Fatal(http.ListenAndServe(":9999", router))
}
