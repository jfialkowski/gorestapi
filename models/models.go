package models

import (
	"database/sql"
)

// Employee Struct
type Employee struct {
	FirstName  string
	LastName   string
	Title      string
	Department string
}

//SelectAllEmployees selects all employees from table and returns []*Employee and nil, or nil and err
func SelectAllEmployees(db *sql.DB) ([]*Employee, error) {

	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allEmployees := make([]*Employee, 0)
	for rows.Next() {
		emp := new(Employee)
		err := rows.Scan(&emp.FirstName, &emp.LastName, &emp.Title, &emp.Department)
		if err != nil {
			return nil, err
		}
		allEmployees = append(allEmployees, emp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return allEmployees, nil
}
