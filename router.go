package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(handlers *Handlers) *mux.Router {
	router := mux.NewRouter()

	// load routes
	for _, route := range GetRoutes(handlers) {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./admin/dist")))

	return router
}
