package main

import (
	"net/http"
)

//Route struct for routing http requests.
type Route struct {
	Name        string
	Method      string
	CertAuth    bool
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes interface
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		false,
		"/",
		home,
	},
	Route{
		"Index",
		"GET",
		false,
		"/v1/",
		home,
	},
	Route{
		"EmployeesIndex",
		"GET",
		true,
		"/v1/employees",
		EmployeesIndex,
	},
	Route{
		"EmployeesInsert",
		"POST",
		true,
		"/v1/employeesinsert",
		EmployeesInsert,
	},
	Route{
		"EmployeesUpdate",
		"PATCH",
		true,
		"/v1/employeesupdate",
		EmployeesUpdate,
	},
	Route{
		"EmployeesDelete",
		"POST",
		true,
		"/v1/employeesdelete",
		EmployeesDelete,
	},
}
