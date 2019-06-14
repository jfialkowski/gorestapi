package main

import (
	"gorestapi/models"
	"log"
	"os"
)

func main() {

	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	//LoadConfig does just that, load your config
	LoadConfig()

	//ConnectDB connects to Database
	models.ConnectDB(DBuser, DBpass, DBhost, DBport, DBname)

	//Start TLS Enabled Web Server
	server := NewServer()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
