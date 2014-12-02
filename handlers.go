package main

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

type readDirResponse struct {
	Data []*rangolib.File `json:"data"`
}

type createDirResponse struct {
	Dir *rangolib.File `json:"dir"`
}

type updateDirResponse struct {
	Dir *rangolib.File `json:"dir"`
}

// readDir reads contents of a directory
func readDir(w http.ResponseWriter, r *http.Request) {
	fp, err := convertPath(mux.Vars(r)["path"])
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

	printJson(w, &readDirResponse{Data: contents})
}

// createDir creates a directory
func createDir(w http.ResponseWriter, r *http.Request) {

	// combine parent and dirname
	parent := mux.Vars(r)["path"]
	dirname := sanitize.Path(r.FormValue("dir[name]"))
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
	printJson(w, &createDirResponse{Dir: dir})
}

// updateDir renames a directory
func updateDir(w http.ResponseWriter, r *http.Request) {
	fp, err := convertPath(mux.Vars(r)["path"])
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
	dirname := sanitize.Path(r.FormValue("dir[name]"))
	dest := filepath.Join(parent, dirname)

	// rename directory
	dir, err := rangolib.UpdateDir(fp, dest)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// print info
	printJson(w, &updateDirResponse{Dir: dir})
}

// destroyDir deletes a directory
func destroyDir(w http.ResponseWriter, r *http.Request) {
	fp, err := convertPath(mux.Vars(r)["path"])
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

type readPageResponse struct {
	Page *rangolib.Page `json:"page"`
}

type createPageResponse struct {
	Page *rangolib.Page `json:"page"`
}

type updatePageResponse struct {
	Page *rangolib.Page `json:"page"`
}

// readPage reads page data
func readPage(w http.ResponseWriter, r *http.Request) {
	fp, err := convertPath(mux.Vars(r)["path"])
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
	printJson(w, &readPageResponse{Page: page})
}

// createPage creates a new page
func createPage(w http.ResponseWriter, r *http.Request) {
	fp, err := convertPath(mux.Vars(r)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// check that parent dir exists
	if fileExists(fp) || dirExists(fp) == false {
		errDirNotFound.Write(w)
		return
	}

	metastring := r.FormValue("page[meta]")
	if len(metastring) == 0 {
		errNoMeta.Write(w)
	}

	metadata := rangolib.Frontmatter{}
	err = json.Unmarshal([]byte(metastring), &metadata)
	if err != nil {
		errInvalidJson.Write(w)
		return
	}

	content := []byte(r.FormValue("page[content]"))

	page, err := rangolib.CreatePage(fp, metadata, content)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, contentDir)

	printJson(w, &createPageResponse{Page: page})
}

// updatePage writes page data to a file
func updatePage(w http.ResponseWriter, r *http.Request) {
	fp, err := convertPath(mux.Vars(r)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// check that existing page exists
	if dirExists(fp) || fileExists(fp) == false {
		errPageNotFound.Write(w)
		return
	}

	metastring := r.FormValue("page[meta]")
	if len(metastring) == 0 {
		errNoMeta.Write(w)
	}

	metadata := rangolib.Frontmatter{}
	err = json.Unmarshal([]byte(metastring), &metadata)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	content := []byte(r.FormValue("page[content]"))

	page, err := rangolib.UpdatePage(fp, metadata, content)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, contentDir)

	printJson(w, &updatePageResponse{Page: page})
}

// destroyPage deletes a page
func destroyPage(w http.ResponseWriter, r *http.Request) {
	fp, err := convertPath(mux.Vars(r)["path"])
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

// readConfig reads data from a config
func readConfig(w http.ResponseWriter, r *http.Request) {
	config, err := rangolib.ReadConfig()
	if err != nil {
		errNoConfig.Write(w)
		return
	}

	printJson(w, config)
}

// updateConfig writes json data to a config file
func updateConfig(w http.ResponseWriter, r *http.Request) {

	// parse the config
	config := &rangolib.Frontmatter{}
	err := json.Unmarshal([]byte(r.FormValue("config")), config)
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

//  ┬ ┬┬ ┬┌─┐┌─┐
//  ├─┤│ ││ ┬│ │
//  ┴ ┴└─┘└─┘└─┘

func publishSite(w http.ResponseWriter, r *http.Request) {
	output, err := rangolib.RunHugo()
	if err != nil {
		wrapError(err).Write(w)
	}

	printJson(w, struct {
		Output string `json:"output"`
	}{
		Output: string(output),
	})
}

//  ┌─┐┬┬  ┌─┐┌─┐
//  ├┤ ││  ├┤ └─┐
//  └  ┴┴─┘└─┘└─┘

// copy copies a page to a new file
func copy(w http.ResponseWriter, r *http.Request) {
	location := r.Header.Get("Content-Location")
	vars := mux.Vars(r)
	if len(location) > 0 {
		fmt.Fprint(w, "Moving file from "+location+" to "+vars["path"])
	}
}
