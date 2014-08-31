package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	r.HandleFunc("/admin", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the admin page")
	})

	r.HandleFunc("/gallery", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the gallery!")
	})

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}
