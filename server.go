package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/stayradiated/rango/rangolib"
)

type catBody struct {
	path string
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/ls", func(w http.ResponseWriter, req *http.Request) {
		files := rangolib.Files()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(files)
	})

	r.HandleFunc("/cat", func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			panic(err)
		}

		path := req.FormValue("path")

		file, err := os.Open("content/" + path)
		if err != nil {
			fmt.Fprint(w, "Could not open file.")
			return
		}
		page, err := rangolib.Read(file)
		if err != nil {
			fmt.Fprint(w, "Could not parse file.")
			return
		}
		json.NewEncoder(w).Encode(*page)
	})

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")

}
