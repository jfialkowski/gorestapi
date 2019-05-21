package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Homepage, nothing really to see here! ")
}

func EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	employees := Employees{
		Employee{FirstName: "John", LastName: "Doe", Title: "Developer", Department: "Information Technology"},
		Employee{FirstName: "Jane", LastName: "Smith", Title: "Manager", Department: "Information Technology"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(employees); err != nil {
		panic(err)
	}
}
