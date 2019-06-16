package main

import (
	"fmt"
	"log"
)

// Employee Struct
type Employee struct {
	FirstName  string
	LastName   string
	Title      string
	Department string
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
		err := rows.Scan(&emp.FirstName, &emp.LastName, &emp.Title, &emp.Department)
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
