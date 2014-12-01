package server

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ReadDir",
		"GET", "/api/dir/{path:.*}", readDir,
	},
	Route{
		"CreateDir",
		"POST", "/api/dir/{path:.*}", createDir,
	},
}
