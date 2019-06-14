package main

import (
	"encoding/json"
	"fmt"
	"gorestapi/models"
	"net/http"
)

type myEnv struct {
	*models.Env
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Homepage, nothing really to see here! ")
}

// EmployeesIndex returns all employees
func (env myEnv) EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	employees, err := models.SelectAllEmployees(*models.DB)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(employees); err != nil {
		panic(err)
	}
}
