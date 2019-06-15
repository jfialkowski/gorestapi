package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gorestapi/models/"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Homepage, nothing really to see here! ")
}

// EmployeesIndex returns all employees
func EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	} else {
		employees, err := models.SelectAllEmployees(db)
		if err != nil {
			log.Fatal(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(employees); err != nil {
			panic(err)
		}
	}
}
