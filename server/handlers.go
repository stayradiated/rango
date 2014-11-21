package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	"github.com/stayradiated/rango/rangolib"
)

//  ┌┬┐┬┬─┐┌─┐┌─┐┌┬┐┌─┐┬─┐┬┌─┐┌─┐
//   │││├┬┘├┤ │   │ │ │├┬┘│├┤ └─┐
//  ─┴┘┴┴└─└─┘└─┘ ┴ └─┘┴└─┴└─┘└─┘

type handleReadDirResponse struct {
	Data []*rangolib.File `json:"data"`
}

type handleCreateDirResponse struct {
	Dir *rangolib.File `json:"dir"`
}

type handleUpdateDirResponse struct {
	Dir *rangolib.File `json:"dir"`
}

// handleReadDir reads contents of a directory
func handleReadDir(w http.ResponseWriter, req *http.Request) {
	fp, err := convertPath(mux.Vars(req)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// try and read contents of dir
	contents, err := rangolib.ReadDir(fp)
	if err != nil {
		errDirNotFound.Write(w)
		return
	}

	// trim content prefix
	for _, item := range contents {
		item.Path = strings.TrimPrefix(item.Path, contentDir)
	}

	printJson(w, &handleReadDirResponse{Data: contents})
}

// handleCreateDir creates a directory
func handleCreateDir(w http.ResponseWriter, req *http.Request) {

	// combine parent and dirname
	parent := mux.Vars(req)["path"]
	dirname := req.FormValue("dir[name]")
	fp := filepath.Join(parent, dirname)

	// check that it is a valid path
	fp, err := convertPath(fp)
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// check if dir already exists
	if fileExists(fp) || dirExists(fp) {
		errDirConflict.Write(w)
		return
	}

	// make directory
	dir, err := rangolib.CreateDir(fp)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix
	dir.Path = strings.TrimPrefix(dir.Path, contentDir)

	// print info
	printJson(w, &handleCreateDirResponse{Dir: dir})
}

// handleUpdateDir renames a directory
func handleUpdateDir(w http.ResponseWriter, req *http.Request) {
	fp, err := convertPath(mux.Vars(req)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// check that the specified directory is not the root content folder
	if fp == contentDir {
		errInvalidDir.Write(w)
		return
	}

	// check that directory exists
	if dirExists(fp) == false {
		errDirNotFound.Write(w)
		return
	}

	// combine parent dir with dir name
	parent := filepath.Dir(fp)
	dirname := sanitize.Path(req.FormValue("dir[name]"))
	dest := filepath.Join(parent, dirname)

	// rename directory
	dir, err := rangolib.UpdateDir(fp, dest)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// print info
	printJson(w, &handleUpdateDirResponse{Dir: dir})
}

// handleDeleteDir deletes a directory
func handleDeleteDir(w http.ResponseWriter, req *http.Request) {
	fp, err := convertPath(mux.Vars(req)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// check that the specified directory is not the root content folder
	if fp == contentDir {
		errInvalidDir.Write(w)
		return
	}

	// remove directory
	if err = rangolib.DeleteDir(fp); err != nil {
		errDirNotFound.Write(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//  ┌─┐┌─┐┌─┐┌─┐┌─┐
//  ├─┘├─┤│ ┬├┤ └─┐
//  ┴  ┴ ┴└─┘└─┘└─┘

type handleReadPageResponse struct {
	Page *rangolib.Page `json:"page"`
}

type handleCreatePageResponse struct {
	Page *rangolib.Page `json:"page"`
}

type handleUpdatePageResponse struct {
	Page *rangolib.Page `json:"page"`
}

// handleReadPage reads page data
func handleReadPage(w http.ResponseWriter, req *http.Request) {
	fp, err := convertPath(mux.Vars(req)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// read page from disk
	page, err := rangolib.ReadPage(fp)
	if err != nil {
		errPageNotFound.Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, contentDir)

	// print json
	printJson(w, &handleReadPageResponse{Page: page})
}

// handleCreatePage creates a new page
func handleCreatePage(w http.ResponseWriter, req *http.Request) {
	fp, err := convertPath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// check that parent dir exists
	if fileExists(fp) || dirExists(fp) == false {
		errDirNotFound.Write(w)
		return
	}

	metastring := req.FormValue("page[meta]")
	if len(metastring) == 0 {
		errNoMeta.Write(w)
	}

	metadata := rangolib.Frontmatter{}
	err = json.Unmarshal([]byte(metastring), &metadata)
	if err != nil {
		errInvalidJson.Write(w)
		return
	}

	content := []byte(req.FormValue("page[content]"))

	page, err := rangolib.CreatePage(fp, metadata, content)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, contentDir)

	printJson(w, &handleCreatePageResponse{Page: page})
}

// handleUpdatePage writes page data to a file
func handleUpdatePage(w http.ResponseWriter, req *http.Request) {
	fp, err := convertPath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// check that existing page exists
	if dirExists(fp) || fileExists(fp) == false {
		errPageNotFound.Write(w)
		return
	}

	metastring := req.FormValue("page[meta]")
	if len(metastring) == 0 {
		errNoMeta.Write(w)
	}

	metadata := rangolib.Frontmatter{}
	err = json.Unmarshal([]byte(metastring), &metadata)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	content := []byte(req.FormValue("page[content]"))

	page, err := rangolib.UpdatePage(fp, metadata, content)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, contentDir)

	printJson(w, &handleUpdatePageResponse{Page: page})
}

// handleDeletePage deletes a page
func handleDeletePage(w http.ResponseWriter, req *http.Request) {
	fp, err := convertPath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// delete page
	if err = rangolib.DeletePage(fp); err != nil {
		errPageNotFound.Write(w)
		return
	}

	// don't need to send anything back
	w.WriteHeader(http.StatusNoContent)
}

//  ┌─┐┌─┐┌┐┌┌─┐┬┌─┐
//  │  │ ││││├┤ ││ ┬
//  └─┘└─┘┘└┘└  ┴└─┘

// handleReadConfig reads data from a config
func handleReadConfig(w http.ResponseWriter, req *http.Request) {
	config, err := rangolib.ReadConfig()
	if err != nil {
		errNoConfig.Write(w)
		return
	}

	printJson(w, config)
}

// handleUpdateConfig writes json data to a config file
func handleUpdateConfig(w http.ResponseWriter, req *http.Request) {

	// parse the config
	config := &rangolib.Frontmatter{}
	err := json.Unmarshal([]byte(req.FormValue("config")), config)
	if err != nil {
		errInvalidJson.Write(w)
		return
	}

	// save config
	if err := rangolib.SaveConfig(config); err != nil {
		wrapError(err).Write(w)
		return
	}

	// don't need to send anything back
	w.WriteHeader(http.StatusNoContent)
}

//  ┌─┐┬┬  ┌─┐┌─┐
//  ├┤ ││  ├┤ └─┐
//  └  ┴┴─┘└─┘└─┘

// handleCopy copies a page to a new file
func handleCopy(w http.ResponseWriter, req *http.Request) {
	location := req.Header.Get("Content-Location")
	vars := mux.Vars(req)
	if len(location) > 0 {
		fmt.Fprint(w, "Moving file from "+location+" to "+vars["path"])
	}
}
