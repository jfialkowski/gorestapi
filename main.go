package main

import (
	"fmt"
	"log"
	"net/http"
)

// Declare struct of Employee Data
type Employee struct {
	firstName   string `json:"firstName"`
	lastName    string `json:"lastName"`
	title       string `json:"title"`
	deptartment string `json:"department"`
}

// Declare a global Employee array to be used in main function
var Employees []Employee

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Homepage, nothing really to see here! ")
	fmt.Println("Endpoint Accessed: home ")
}

func handleRequests() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func main() {

	//Some dummy Data

	employees := Employees{
		Employee{firstName: "John", lastName: "Doe", title: "Senior Developer", department: "Information Technology"},
		Employee{firstName: "Jane", lastName: "Smith", title: "Manager", department: "Information Technology"},
	}
	handleRequests()
}
