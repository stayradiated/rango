package server

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
	fp, err := sanitizePath(mux.Vars(req)["path"])
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
	parent, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		errInvalidDir.Write(w)
		return
	}

	// combine parent dir with dir name
	dirname := req.FormValue("dir[name]")
	fp := filepath.Join(parent, dirname)

	// check if dir already exists
	if dirExists(fp) {
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
	fp, err := sanitizePath(mux.Vars(req)["path"])
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
	fp, err := sanitizePath(mux.Vars(req)["path"])
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

	// remove directory
	if err = rangolib.DeleteDir(fp); err != nil {
		wrapError(err).Write(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
