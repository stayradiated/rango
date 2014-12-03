package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	"github.com/stayradiated/rango/rangolib"
)

type Handlers struct {
	Config     rangolib.ConfigManager
	Dir        rangolib.DirManager
	Page       rangolib.PageManager
	ContentDir string
}

//  ┌┬┐┬┬─┐┌─┐┌─┐┌┬┐┌─┐┬─┐┬┌─┐┌─┐
//   │││├┬┘├┤ │   │ │ │├┬┘│├┤ └─┐
//  ─┴┘┴┴└─└─┘└─┘ ┴ └─┘┴└─┴└─┘└─┘

type readDirResponse struct {
	Data rangolib.Files `json:"data"`
}

type createDirResponse struct {
	Dir *rangolib.File `json:"dir"`
}

type updateDirResponse struct {
	Dir *rangolib.File `json:"dir"`
}

// readDir reads contents of a directory
func (h Handlers) ReadDir(w http.ResponseWriter, r *http.Request) {
	fp, err := h.convertPath(mux.Vars(r)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// try and read contents of dir
	var contents rangolib.Files
	contents, err = h.Dir.Read(fp)
	if err != nil {
		errDirNotFound.Write(w)
		return
	}

	// trim content prefix
	for _, item := range contents {
		item.Path = strings.TrimPrefix(item.Path, h.ContentDir)
	}

	printJson(w, &readDirResponse{Data: contents})
}

// createDir creates a directory
func (h Handlers) CreateDir(w http.ResponseWriter, r *http.Request) {

	// combine parent and dirname
	parent := mux.Vars(r)["path"]
	dirname := sanitize.Path(r.FormValue("dir[name]"))
	fp := filepath.Join(parent, dirname)

	// check that it is a valid path
	fp, err := h.convertPath(fp)
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
	dir, err := h.Dir.Create(fp)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix
	dir.Path = strings.TrimPrefix(dir.Path, h.ContentDir)

	// print info
	printJson(w, &createDirResponse{Dir: dir})
}

// updateDir renames a directory
func (h Handlers) UpdateDir(w http.ResponseWriter, r *http.Request) {
	fp, err := h.convertPath(mux.Vars(r)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// check that the specified directory is not the root content folder
	if fp == h.ContentDir {
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
	dir, err := h.Dir.Update(fp, dest)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// print info
	printJson(w, &updateDirResponse{Dir: dir})
}

// destroyDir deletes a directory
func (h Handlers) DestroyDir(w http.ResponseWriter, r *http.Request) {
	fp, err := h.convertPath(mux.Vars(r)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// check that the specified directory is not the root content folder
	if fp == h.ContentDir {
		errInvalidDir.Write(w)
		return
	}

	// remove directory
	if err = h.Dir.Destroy(fp); err != nil {
		errDirNotFound.Write(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//  ┌─┐┌─┐┌─┐┌─┐┌─┐
//  ├─┘├─┤│ ┬├┤ └─┐
//  ┴  ┴ ┴└─┘└─┘└─┘

type readPageResponse struct {
	Page *rangolib.PageFile `json:"page"`
}

type createPageResponse struct {
	Page *rangolib.PageFile `json:"page"`
}

type updatePageResponse struct {
	Page *rangolib.PageFile `json:"page"`
}

// readPage reads page data
func (h Handlers) ReadPage(w http.ResponseWriter, r *http.Request) {
	fp, err := h.convertPath(mux.Vars(r)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// read page from disk
	page, err := h.Page.Read(fp)
	if err != nil {
		errPageNotFound.Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, h.ContentDir)

	// print json
	printJson(w, &readPageResponse{Page: page})
}

// createPage creates a new page
func (h Handlers) CreatePage(w http.ResponseWriter, r *http.Request) {
	fp, err := h.convertPath(mux.Vars(r)["path"])
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

	page, err := h.Page.Create(fp, metadata, content)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, h.ContentDir)

	printJson(w, &createPageResponse{Page: page})
}

// updatePage writes page data to a file
func (h Handlers) UpdatePage(w http.ResponseWriter, r *http.Request) {
	fp, err := h.convertPath(mux.Vars(r)["path"])
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

	page, err := h.Page.Update(fp, metadata, content)
	if err != nil {
		wrapError(err).Write(w)
		return
	}

	// trim content prefix from path
	page.Path = strings.TrimPrefix(page.Path, h.ContentDir)

	printJson(w, &updatePageResponse{Page: page})
}

// destroyPage deletes a page
func (h Handlers) DestroyPage(w http.ResponseWriter, r *http.Request) {
	fp, err := h.convertPath(mux.Vars(r)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// delete page
	if err = h.Page.Destroy(fp); err != nil {
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
func (h Handlers) ReadConfig(w http.ResponseWriter, r *http.Request) {
	config, err := h.Config.Parse()
	if err != nil {
		errNoConfig.Write(w)
		return
	}

	printJson(w, config)
}

// updateConfig writes json data to a config file
func (h Handlers) UpdateConfig(w http.ResponseWriter, r *http.Request) {

	// parse the config
	config := &rangolib.ConfigMap{}
	err := json.Unmarshal([]byte(r.FormValue("config")), config)
	if err != nil {
		errInvalidJson.Write(w)
		return
	}

	// save config
	if err := h.Config.Save(config); err != nil {
		wrapError(err).Write(w)
		return
	}

	// don't need to send anything back
	w.WriteHeader(http.StatusNoContent)
}

//  ┬ ┬┬ ┬┌─┐┌─┐
//  ├─┤│ ││ ┬│ │
//  ┴ ┴└─┘└─┘└─┘

func (h Handlers) PublishSite(w http.ResponseWriter, r *http.Request) {
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

func (h Handlers) convertPath(p string) (string, error) {
	err := errors.New("invalid path")

	// join path with content folder
	fp := path.Join(h.ContentDir, p)

	// check that path still starts with content dir
	if !strings.HasPrefix(fp, h.ContentDir) {
		return fp, err
	}

	// check that path doesn't contain any ..
	if strings.Contains(fp, "..") {
		return fp, err
	}

	return fp, nil
}
