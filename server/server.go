package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var contentDir = "content"

// main starts the web server
func Start() *negroni.Negroni {

	// setup middleware
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewStatic(http.Dir("./admin/dist")))
	n.Use(negroni.HandlerFunc(respondWithJson))

	// listen to handlers
	n.UseHandler(Handlers())

	// return negroni
	return n
}

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

// respondWithJson sets the content-type header to application/json
func respondWithJson(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, req)
}
