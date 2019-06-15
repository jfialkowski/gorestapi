package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	db *sql.DB
}

//ConnectDB connects to a DB
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", DBuser+":"+DBpass+"@tcp("+DBhost+":"+DBport+")/"+DBname+"?tls=skip-verify&autocommit=true")
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		return nil, err
	} else {
		log.Println("Connected to Database")
	}

	return db, nil
}
