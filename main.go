package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	router := NewRouter()
	LoadConfig()
	ConnectDB(DBuser, DBpass, DBhost, DBport, DBname)
	log.Fatal(http.ListenAndServe(":9999", router))
}
