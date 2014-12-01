package main

import "github.com/gorilla/mux"

// Handlers connects the handlers to a new router
func Handlers() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	// directories
	api.HandleFunc("/dir/{path:.*}", handleReadDir).Methods("GET")
	api.HandleFunc("/dir/{path:.*}", handleCreateDir).Methods("POST")
	api.HandleFunc("/dir/{path:.*}", handleUpdateDir).Methods("PUT")
	api.HandleFunc("/dir/{path:.*}", handleDeleteDir).Methods("DELETE")

	// pages
	api.HandleFunc("/page/{path:.*}", handleReadPage).Methods("GET")
	api.HandleFunc("/page/{path:.*}", handleCreatePage).Methods("POST")
	api.HandleFunc("/page/{path:.*}", handleUpdatePage).Methods("PUT")
	api.HandleFunc("/page/{path:.*}", handleDeletePage).Methods("DELETE")

	// config
	api.HandleFunc("/config", handleReadConfig).Methods("GET")
	api.HandleFunc("/config", handleUpdateConfig).Methods("PUT")

	// hugo
	api.HandleFunc("/site/publish", handlePublishSite).Methods("POST")

	// files
	api.HandleFunc("/copy/{path:.*}", handleCopy).Methods("POST")

	return r
}
