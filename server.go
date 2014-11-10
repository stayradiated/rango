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

	r.HandleFunc("/config", getConfig).Methods("GET")
	r.HandleFunc("/config", setConfig).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")

}

func getPages(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(req.FormValue("path"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	pathList, err := rangolib.Files(fp)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	writeJson(w, pathList)
}

func getPage(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(req.FormValue("path"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	file, err := os.Open(fp)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	page, err := rangolib.Read(file)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	writeJson(w, page)
}

func setPage(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(req.FormValue("path"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	metadata := map[string]interface{}{}
	err = json.Unmarshal([]byte(req.FormValue("metadata")), &metadata)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	content := []byte(req.FormValue("content"))

	rangolib.Save(fp, metadata, content)
}

func getConfig(w http.ResponseWriter, req *http.Request) {
	config, err := rangolib.ReadConfig()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	writeJson(w, config)
}

func setConfig(w http.ResponseWriter, req *http.Request) {
	config := map[string]interface{}{}

	if err := json.Unmarshal([]byte(req.FormValue("config")), &config); err != nil {
		fmt.Fprint(w, err)
		return
	}

	if err := rangolib.SaveConfig(config); err != nil {
		fmt.Fprint(w, err)
		return
	}
}

/* HELPERS */

func sanitizePath(p string) (string, error) {
	fp := path.Join("content", p)

	if !strings.HasPrefix(fp, "content") || strings.Contains(fp, "..") {
		return fp, errors.New("Invalid Path")
	}

	return fp, nil
}

func writeJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(obj)
}
