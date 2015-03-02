package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterConfig struct {
	Handlers  *Handlers
	AdminDir  string
	AssetsDir string
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

	// serve static assets (user images, etc)
	assetsFs := http.FileServer(http.Dir(config.AssetsDir))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", assetsFs))

	// serve admin client files (html, css, etc)
	adminFs := http.FileServer(http.Dir(config.AdminDir))
	router.PathPrefix("/").Handler(adminFs)

	return router
}
