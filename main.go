package main

import (
	"log"
	"os"
)

func main() {

	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	//LoadConfig does just that, load your config
	LoadConfig()
	var err error
	//ConnectDB connects to Database
	DBCon, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	//Start TLS Enabled Web Server
	server := NewServer()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
