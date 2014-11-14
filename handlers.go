package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	"github.com/stayradiated/rango/rangolib"
)

/*
 *     __   __   ___      ___  ___
 *    /  ` |__) |__   /\   |  |__
 *    \__, |  \ |___ /~~\  |  |___
 *
 */

// handleCreateDir creates a directory
func handleCreateDir(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	os.MkdirAll(fp, 0755)

	returnSuccess(w, map[string]interface{}{
		"success": true,
		"path":    fp,
	})
}

func handleCreateFile(w http.ResponseWriter, req *http.Request) {
}

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
	fp = path.Join(fp, title)

	content := []byte(req.FormValue("content"))

	rangolib.Save(fp, metadata, content)

	returnSuccess(w, map[string]interface{}{
		"success": true,
		"path":    fp,
	})
}

// handleCopy copies a page to a new file
func handleCopy(w http.ResponseWriter, req *http.Request) {
	location := req.Header.Get("Content-Location")
	vars := mux.Vars(req)
	if len(location) > 0 {
		fmt.Fprint(w, "Moving file from "+location+" to "+vars["path"])
	}
}

/*
 *     __   ___       __
 *    |__) |__   /\  |  \
 *    |  \ |___ /~~\ |__/
 *
 */

// handleReadDir reads contents of a directory
func handleReadDir(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	pathList, err := rangolib.Files(fp)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	returnSuccess(w, pathList)
}

// handleReadFile
func handleReadFile(w http.ResponseWriter, req *http.Request) {
}

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

	returnSuccess(w, page)
}

// handleReadConfig reads data from a config
func handleReadConfig(w http.ResponseWriter, req *http.Request) {
	config, err := rangolib.ReadConfig()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	returnSuccess(w, config)
}

/*
 *          __   __       ___  ___
 *    |  | |__) |  \  /\   |  |__
 *    \__/ |    |__/ /~~\  |  |___
 *
 */

func handleRename(w http.ResponseWriter, req *http.Request) {
}

func handleUpdateFile(w http.ResponseWriter, req *http.Request) {
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
	fp = path.Join(fp, title)

	content := []byte(req.FormValue("content"))

	rangolib.Save(fp, metadata, content)
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

/*
 *     __   ___  __  ___  __   __
 *    |  \ |__  /__`  |  |__) /  \ \ /
 *    |__/ |___ .__/  |  |  \ \__/  |
 *
 */

// handleDestroy deletes a file, page or directory
func handleDestroy(w http.ResponseWriter, req *http.Request) {
	fp, err := sanitizePath(mux.Vars(req)["path"])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if err = os.RemoveAll(fp); err != nil {
		fmt.Fprint(w, err)
		return
	}

	returnSuccess(w, map[string]interface{}{
		"path":    fp,
		"success": true,
	})
}

/*
 *         ___         __
 *    |  |  |  | |    /__`
 *    \__/  |  | |___ .__/
 *
 */

func sanitizePath(p string) (string, error) {
	fp := path.Join("content", p)

	if !strings.HasPrefix(fp, "content") || strings.Contains(fp, "..") {
		return fp, errors.New("Invalid Path")
	}

	return fp, nil
}

func returnSuccess(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
