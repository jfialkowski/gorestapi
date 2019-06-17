package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Homepage, nothing really to see here! ")
}

// EmployeesInsert handler for creating a new employee record
func EmployeesInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}
	Emp := Employee{}
	err = json.Unmarshal(jsn, &Emp)
	if err != nil {
		log.Println("Decoding error: ", err)
	}
	result, err := InsertEmployee(Emp)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(405), 405)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, result)
}

//EmployeesUpdate handler for updating employee
func EmployeesUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}
	Emp := Employee{}
	err = json.Unmarshal(jsn, &Emp)
	if err != nil {
		log.Println("Decoding error: ", err)
	}
	result, err := UpdateEmployee(Emp)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(405), 405)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, result)
}

//EmployeesDelete handler for deleting employees
func EmployeesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}
	Emp := Employee{}
	err = json.Unmarshal(jsn, &Emp)
	if err != nil {
		log.Println("Decoding error: ", err)
	}
	fmt.Printf("Emp being sent over is: %v", &Emp)
	result, err := DeleteEmployee(&Emp)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(405), 405)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, result)
}

// EmployeesIndex handler returns all employees
func EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	employees, err := SelectAllEmployees()
	if err != nil {
		log.Fatal(err)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		js, err := json.Marshal(employees)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(js))
	}
}
