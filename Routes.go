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
		"EmployeesIndex",
		"GET",
		true,
		"/employees",
		EmployeesIndex,
	},
	Route{
		"EmployeesInsert",
		"POST",
		true,
		"/employeesinsert",
		EmployeesInsert,
	},
	Route{
		"EmployeesUpdate",
		"PATCH",
		true,
		"/employeesupdate",
		EmployeesUpdate,
	},
	Route{
		"EmployeesDelete",
		"POST",
		true,
		"/employeesdelete",
		EmployeesDelete,
	},
}
