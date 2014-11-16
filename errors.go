package main

import "net/http"

type apiError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
}

func NewApiError(err error) *apiError {
	return &apiError{
		Status: http.StatusInternalServerError,
		Code:   "internal_server_error",
		Title:  err.Error(),
	}
}

func (err *apiError) Write(w http.ResponseWriter) {
	w.WriteHeader(err.Status)
	printError(w, err)
}

var errInvalidDir = &apiError{
	Status: http.StatusBadRequest,
	Code:   "invalid_dir",
	Title:  "invalid directory",
}

var errDirNotFound = &apiError{
	Status: http.StatusNotFound,
	Code:   "dir_not_found",
	Title:  "directory does not exist",
}

var errMalformedJson = &apiError{
	Status: http.StatusBadRequest,
	Code:   "malformed_json",
	Title:  "could not parse request body",
}

var errDirConflict = &apiError{
	Status: http.StatusConflict,
	Code:   "dir_conflict",
	Title:  "directory already exists",
}
