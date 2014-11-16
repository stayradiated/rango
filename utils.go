package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

func dirExists(fp string) bool {
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func sanitizePath(p string) (string, error) {
	err := errors.New("invalid path")

	// join path with content folder
	fp := path.Join(CONTENT_DIR, p)

	// check that path still starts with content dir
	if !strings.HasPrefix(fp, CONTENT_DIR) {
		return fp, err
	}

	// check that path doesn't contain any ..
	if strings.Contains(fp, "..") {
		return fp, err
	}

	return fp, nil
}

func printError(w io.Writer, err interface{}) {
	printJson(w, struct {
		Errors interface{} `json:"errors"`
	}{
		Errors: err,
	})
}

func printJson(w io.Writer, obj interface{}) {
	json.NewEncoder(w).Encode(obj)
}
