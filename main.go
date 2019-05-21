package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Employee Struct
type Employee struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Title      string `json:"title"`
	Department string `json:"department"`
}

type Employees []Employee

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Homepage, nothing really to see here! ")
	fmt.Println("Endpoint Accessed: home ")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/employees", EmployeesIndex)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	employees := Employees{
		Employee{FirstName: "John", LastName: "Doe", Title: "Senior Developer", Department: "Information Technology"},
		Employee{FirstName: "Jane", LastName: "Smith", Title: "Manager", Department: "Information Technology"},
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {
	handleRequests()
}
