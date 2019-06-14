package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	db *sql.DB
}

//ConnectDB connects to a DB
func ConnectDB(username string, password string, host string, port string, dbname string) (*sql.DB, error) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?tls=skip-verify&autocommit=true")
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
