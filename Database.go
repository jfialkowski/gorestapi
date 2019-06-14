package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBCon *sql.DB
)

//ConnectDB connects to a DB
func ConnectDB(username string, password string, host string, port string, dbname string) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?tls=skip-verify&autocommit=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("Ping: " + host + ":" + port + " SUCESS!")
	}

}

func insert() {

}

//SelectAllEmployees selects all employees from table and prints them for now. Should turn into a JSON String or something
func SelectAllEmployees() {

	var (
		firstname  string
		lastname   string
		title      string
		department string
	)

	db, err := sql.Open("mysql", DBuser+":"+DBpass+"@tcp("+DBhost+":"+DBport+")/"+DBname+"?tls=skip-verify&autocommit=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&firstname, &lastname, &title, &department)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(firstname, lastname, title, department)
	}
}

func delete() {

}
