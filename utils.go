package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path"
	"strings"
)

func fileExists(fp string) bool {
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}
	return info.IsDir() == false
}

func dirExists(fp string) bool {
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func convertPath(p string) (string, error) {
	err := errors.New("invalid path")

	// join path with content folder
	fp := path.Join(contentDir, p)

	// check that path still starts with content dir
	if !strings.HasPrefix(fp, contentDir) {
		return fp, err
	}

	// check that path doesn't contain any ..
	if strings.Contains(fp, "..") {
		return fp, err
	}

	return fp, nil
}

func printError(w http.ResponseWriter, err interface{}) {
	printJson(w, err)
}

func printJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(obj)
}
