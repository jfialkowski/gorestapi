package main

// Employee Struct
type Employee struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Title      string `json:"title"`
	Department string `json:"department"`
}

type Employees []Employee
