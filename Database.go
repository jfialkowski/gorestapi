package main

import (
	"database/sql"
	"fmt"

	//"github.com/eduardbcom/gonfig"
	_ "github.com/go-sql-driver/mysql"
)

// type Config struct {
// 	DbConfig struct {
// 		Host         string `json:"host"`
// 		DatabaseName string `json:databasename`
// 		Port         string `json:"port"`
// 		Username     string `json:username`
// 		Password     string `json:password`
// 	} `json:"dbConfig"`
// 	Name string `json:"name"`
// }

func connectDB(username string, password string, host string, port string, dbname string) {
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

// func selectQ() {

// 	appConfig := &Config{}
// 	if rawData, err := gonfig.Read(); err != nil {
// 		panic(err)
// 	} else {
// 		if err := json.Unmarshal(rawData, appConfig); err != nil {
// 			panic(err)
// 		} else {
// 			fmt.Printf(
// 				"{\"name\": \"%s\", \"dbConfig\": {\"host\": \"%s\",\"dbname\": \"%s\", port: \"%d\", \"username\": \"%s\"}}\n",
// 				appConfig.Name,
// 				appConfig.DbConfig.Host,
// 				appConfig.DbConfig.DatabaseName,
// 				appConfig.DbConfig.Port,
// 				appConfig.DbConfig.Username,
// 			) // {"name": "new-awesome-name", "dbConfig": {"host": "prod-db-server", port: "1"}}
// 		}
// 	}
// 	var (
// 		firstname  string
// 		lastname   string
// 		title      string
// 		department string
// 	)
// 	//fmt.Println(appConfig.DbConfig.Username + ":" + appConfig.DbConfig.Password + "@tcp(" + appConfig.DbConfig.Host + ":" + string(appConfig.DbConfig.Port) + ")/" + appConfig.DbConfig.DatabaseName + "?tls=skip-verify&autocommit=true")
// 	db, err := sql.Open("mysql", appConfig.DbConfig.Username+":"+appConfig.DbConfig.Password+"@tcp("+appConfig.DbConfig.Host+":"+appConfig.DbConfig.Port+")/"+appConfig.DbConfig.DatabaseName+"?tls=skip-verify&autocommit=true")
// 	if err != nil {
// 		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM employees")
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan(&firstname, &lastname, &title, &department)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(firstname, lastname, title, department)
// 	}
// }

func delete() {

}

//func main() {

//	selectQ()

//}
