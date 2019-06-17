package main

import (
	"fmt"
	"log"
)

// Employee Struct
type Employee struct {
	EmpID      int
	FirstName  string
	LastName   string
	Title      string
	Department string
}

//InsertEmployee inserts a new Employee record.
func InsertEmployee(FirstName string, LastName string, Title string, Department string) (string, error) {

	result := ""
	stmtIns, err := DBCon.Prepare("INSERT INTO employees (firstname, lastname, title, department) VALUES (?, ?, ?, ?)") // ? = placeholder
	if err != nil {
		log.Println(err) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(FirstName, LastName, Title, Department)
	if err != nil {
		log.Println("Could not insert record")
		result = "{'Status': 'NOK-FAILURE INSERTING DATA'}"
	}
	result = "{'Status': 'OK-SUCESS'}"
	return result, err
}

//SelectAllEmployees selects all employees from table and returns []*Employee and nil, or nil and err
func SelectAllEmployees() ([]*Employee, error) {

	rows, err := DBCon.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allEmployees := make([]*Employee, 0)
	for rows.Next() {
		//fmt.Printf("%v", rows)
		emp := new(Employee)
		err := rows.Scan(&emp.EmpID, &emp.FirstName, &emp.LastName, &emp.Title, &emp.Department)
		fmt.Printf("%v", emp)
		if err != nil {
			log.Println("Nothing returned in rows")
			return nil, err
		}
		allEmployees = append(allEmployees, emp)
	}
	if err = rows.Err(); err != nil {
		log.Println("returning nothing")
		return nil, err
	}
	return allEmployees, nil
}
