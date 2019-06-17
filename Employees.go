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

//UpdateEmployee updates an existing record in DB
func UpdateEmployee(emp Employee) (string, error) {

	result := ""
	stmtIns, err := DBCon.Prepare("UPDATE employees SET (firstname, lastname, title, department) VALUES (?, ?, ?, ?) WHERE empid = ?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmtIns.Exec(emp.FirstName, emp.LastName, emp.Title, emp.Department, emp.EmpID)
	if err != nil {
		log.Println("Could not insert record")
		result = "{'Status': 'NOK-FAILURE UPDATING RECORD'}"
	}
	result = "{'Status': 'OK-SUCCESS'}"
	stmtIns.Close()
	return result, err
}

//DeleteEmployee deletes a record from the database
func DeleteEmployee(emp Employee) (string, error) {
	result := ""
	stmtIns, err := DBCon.Prepare("DELETE from employees WHERE empid = ?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmtIns.Exec(emp.EmpID)
	if err != nil {
		log.Println("Could not delete record")
		result = "{'Status': 'NOK-FAILURE DELETING RECORD'}"
	}
	result = "{'Status': 'OK-SUCCESS'}"
	stmtIns.Close()
	return result, err
}

//InsertEmployee inserts a new Employee record.
func InsertEmployee(emp Employee) (string, error) {
	result := ""
	stmtIns, err := DBCon.Prepare("INSERT INTO employees (firstname, lastname, title, department) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmtIns.Exec(emp.FirstName, emp.LastName, emp.Title, emp.Department)
	if err != nil {
		log.Println("Could not insert record")

		result = "{'Status': 'NOK-FAILURE INSERTING DATA'}"
	}
	result = "{'Status': 'OK-SUCCESS'}"
	stmtIns.Close()
	return result, err
}

//SelectAllEmployees selects all employees from table and returns []*Employee and nil, or nil and err
func SelectAllEmployees() ([]*Employee, error) {
	rows, err := DBCon.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}
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
		rows.Close()
		return nil, err
	}
	rows.Close()
	return allEmployees, nil
}
