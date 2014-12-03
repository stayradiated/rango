package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(handlers *Handlers) *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// load routes
	for _, route := range GetRoutes(handlers) {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		apiRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./admin/dist")))

	return router
}
