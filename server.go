package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	// negroni + mux
	n := negroni.New()
	r := mux.NewRouter()

	// setup middleware
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewStatic(http.Dir("./client/dist/")))
	n.UseHandler(r)

	// API prefix
	api := r.PathPrefix("/api").Subrouter()

	// POST (create)
	api.HandleFunc("/dir/{path:.*}", handleCreateDir).Methods("POST")
	// api.HandleFunc("/file/{path:.*}", handleCreateFile).Methods("POST")
	api.HandleFunc("/page/{path:.*}", handleCreatePage).Methods("POST")
	// api.HandleFunc("/copy/{path:.*}", handleCopy).Methods("POST")

	// GET (read)
	api.HandleFunc("/dir/{path:.*}", handleReadDir).Methods("GET")
	// api.HandleFunc("/file/{path:.*}", handleReadFile).Methods("GET")
	api.HandleFunc("/page/{path:.*}", handleReadPage).Methods("GET")
	api.HandleFunc("/config", handleReadConfig).Methods("GET")

	// PUT (update)
	// api.HandleFunc("/rename/{path:.*}", handleRename).Methods("PUT")
	// api.HandleFunc("/file/{path:.*}", handleUpdateFile).Methods("PUT")
	api.HandleFunc("/page/{path:.*}", handleUpdatePage).Methods("PUT")
	api.HandleFunc("/config", handleUpdateConfig).Methods("PUT")

	// DELETE (destroy)
	api.HandleFunc("/{path:.*}", handleDestroy).Methods("DELETE")

	// start server
	n.Run(":8080")
}
