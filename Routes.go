package main

import (
	"net/http"
)

//Route struct for routing http requests.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes interface
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		home,
	},
	Route{
		"EmployeesIndex",
		"GET",
		"/employees",
		EmployeesIndex,
	},
	Route{
		"Employees",
		"POST",
		"/employees",
		EmployeesIndex,
	},
}
