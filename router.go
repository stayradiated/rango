package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterConfig struct {
	Handlers *Handlers
	AdminDir string
}

func NewRouter(config *RouterConfig) *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// load routes
	for _, route := range GetRoutes(config.Handlers) {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		apiRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(config.AdminDir)))

	return router
}
