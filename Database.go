package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)

//ConnectDB connects to a DB
func ConnectDB() (*sql.DB, error) {
	DBCon, err := sql.Open("mysql", DBuser+":"+DBpass+"@tcp("+DBhost+":"+DBport+")/"+DBname+"?tls=skip-verify&autocommit=true")
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = DBCon.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to Database")

	return DBCon, nil
}
