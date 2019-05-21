package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Employee Struct
type Employee struct {
	firstName  string `json:"firstName"`
	lastName   string `json:"lastName"`
	title      string `json:"title"`
	department string `json:"department"`
}

var Employees []Employee

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Homepage, nothing really to see here! ")
	fmt.Println("Endpoint Accessed: home ")
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Accessed: returnAllEmployees")
	json.NewEncoder(w).Encode(Employees)
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/employees", returnAllEmployees)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func main() {
	//Employees = []Employee{}
	//Employees array to be used in main function
	Employees = append(Employees, Employee{firstName: "John", lastName: "Doe", title: "Senior Developer", department: "Information Technology"})
	Employees = append(Employees, Employee{firstName: "Jane", lastName: "Smith", title: "Manager", department: "Information Technology"})

	handleRequests()
}
