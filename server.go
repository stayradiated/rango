package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/stayradiated/rango/rangolib"
)

type catBody struct {
	path string
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/pages", getPages).Methods("GET")
	r.HandleFunc("/page", getPage).Methods("GET")
	r.HandleFunc("/page", setPage).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")

}

func sanitizePath(p string) (string, error) {
	fp := path.Join("content", p)

	if !strings.HasPrefix(fp, "content") || strings.Contains(fp, "..") {
		return fp, errors.New("Invalid Path")
	}

	return fp, nil
}

func getPages(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(req.FormValue("path"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	pathList, err := rangolib.Files(fp)
	if err != nil {
		fmt.Fprint(w, "ERROR: Could not list folder contents")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(pathList)
}

func getPage(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(req.FormValue("path"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	file, err := os.Open(fp)
	if err != nil {
		fmt.Fprint(w, "Could not open file.")
		return
	}

	page, err := rangolib.Read(file)
	if err != nil {
		fmt.Fprint(w, "Could not parse file.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(*page)
}

func setPage(w http.ResponseWriter, req *http.Request) {
	_, err := sanitizePath(req.FormValue("path"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	metadata := map[string]interface{}{}
	json.Unmarshal([]byte(req.FormValue("metadata")), &metadata)

	fmt.Fprint(w, metadata)
	fmt.Fprint(w, req.FormValue("content"))
}
