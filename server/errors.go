package server

import "net/http"

type apiError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func newApiError(code int, message string) *apiError {
	return &apiError{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	}
}

func wrapError(err error) *apiError {
	return newApiError(http.StatusInternalServerError,
		err.Error())
}

func (a *apiError) Write(w http.ResponseWriter) {
	w.WriteHeader(a.Code)
	printError(w, a)
}

var errInvalidDir = newApiError(http.StatusBadRequest,
	"Invalid Directory")

var errDirNotFound = newApiError(http.StatusNotFound,
	"Directory does not exist")

var errPageNotFound = newApiError(http.StatusNotFound,
	"Page does not exist")

var errMalformedJson = newApiError(http.StatusBadRequest,
	"Could not parse request body")

var errDirConflict = newApiError(http.StatusConflict,
	"Directory already exists")

var errNoMeta = newApiError(http.StatusBadRequest,
	"page[meta] not sent in request")

var errInvalidJson = newApiError(http.StatusBadRequest,
	"Malformed JSON")

var errNoConfig = newApiError(http.StatusNotFound,
	"Could not find config file")
