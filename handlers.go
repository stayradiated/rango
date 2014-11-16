package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	"github.com/stayradiated/rango/rangolib"
)

//   __     __   ___  __  ___  __   __     ___  __
//  |  \ | |__) |__  /  `  |  /  \ |__) | |__  /__`
//  |__/ | |  \ |___ \__,  |  \__/ |  \ | |___ .__/
//

// handleReadDir reads contents of a directory
func handleReadDir(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// check that directory exists
	if dirExists(fp) {
		errDirNotFound.Write(w)
		return
	}

	// try and read contents of dir
	contents, err := rangolib.DirContents(fp)
	if err != nil {
		errDirNotFound.Write(w)
		return
	}

	printJson(w, struct {
		Data []*rangolib.File `json:"data"`
	}{
		Data: contents,
	})
}

// handleCreateDir creates a directory
func handleCreateDir(w http.ResponseWriter, req *http.Request) {
	parent, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	dirname := req.FormValue("dir[name]")
	fp := filepath.Join(parent, dirname)

	if dirExists(fp) {
		errDirConflict.Write(w)
		return
	}

	if err = os.MkdirAll(fp, 0755); err != nil {
		NewApiError(err).Write(w)
		return
	}

	info, err := os.Stat(fp)
	if err != nil {
		NewApiError(err).Write(w)
		return
	}

	dir := &rangolib.File{
		Path: strings.TrimPrefix(fp, CONTENT_DIR),
	}
	dir.Load(info)

	printJson(w, struct {
		Dir *rangolib.File `json:"dir"`
	}{
		Dir: dir,
	})
}

// handleUpdateDir renames a directory
func handleUpdateDir(w http.ResponseWriter, req *http.Request) {
}

// handleDeleteDir deletes a directory
func handleDeleteDir(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if err = os.RemoveAll(fp); err != nil {
		fmt.Fprint(w, err)
		return
	}

	printJson(w, map[string]interface{}{
		"path": fp,
	})
}

//   __        __   ___  __
//  |__)  /\  / _` |__  /__`
//  |    /~~\ \__> |___ .__/
//

// handleReadPage reads page data
func handleReadPage(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
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

	printJson(w, page)
}

// handleCreatePage creates a new page
func handleCreatePage(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
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

	rawTitle, ok := metadata["title"]
	if !ok {
		fmt.Fprint(w, errors.New("Must specify title in metadata"))
		return
	}
	title := sanitize.Path(rawTitle.(string) + ".md")
	fp = filepath.Join(fp, title)

	content := []byte(req.FormValue("content"))

	rangolib.Save(fp, metadata, content)

	printJson(w, map[string]interface{}{
		"path": fp,
	})
}

// handleUpdatePage writes page data to a file
func handleUpdatePage(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
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

	rawTitle, ok := metadata["title"]
	if !ok {
		fmt.Fprint(w, errors.New("Must specify title in metadata"))
		return
	}
	title := sanitize.Path(rawTitle.(string) + ".md")
	fp = filepath.Join(fp, title)

	content := []byte(req.FormValue("content"))

	rangolib.Save(fp, metadata, content)
}

// handleDeletePage deletes a page
func handleDeletePage(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if err = os.Remove(fp); err != nil {
		fmt.Fprint(w, err)
		return
	}

	printJson(w, map[string]interface{}{
		"path": fp,
	})
}

//   __   __        ___    __
//  /  ` /  \ |\ | |__  | / _`
//  \__, \__/ | \| |    | \__>
//

// handleReadConfig reads data from a config
func handleReadConfig(w http.ResponseWriter, req *http.Request) {
	config, err := rangolib.ReadConfig()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	printJson(w, config)
}

// handleUpdateConfig writes json data to a config file
func handleUpdateConfig(w http.ResponseWriter, req *http.Request) {
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

//   ___         ___  __
//  |__  | |    |__  /__`
//  |    | |___ |___ .__/
//

// handleCopy copies a page to a new file
func handleCopy(w http.ResponseWriter, req *http.Request) {
	location := req.Header.Get("Content-Location")
	vars := mux.Vars(req)
	if len(location) > 0 {
		fmt.Fprint(w, "Moving file from "+location+" to "+vars["path"])
	}
}
