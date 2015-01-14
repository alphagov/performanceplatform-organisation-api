package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

// ErrorInfo is as described at jsonapi.org
type ErrorInfo struct {
	ID     string   `json:"id,omitempty"`
	HREF   string   `json:"href,omitempty"`
	Status string   `json:"status,omitempty"`
	Code   string   `json:"code,omitempty"`
	Title  string   `json:"title,omitempty"`
	Detail string   `json:"detail,omitempty"`
	Links  []string `json:"links,omitempty"`
	Path   string   `json:"path,omitempty"`
}

// APIResponse is used for all JSON API responses.
// See jsonapi.org/format/
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  []ErrorInfo `json:"errors"`
}

var (
	renderer = render.New(render.Options{})
)

func main() {
	http.ListenAndServe(":3060", newHandler())
}

func newHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/_status", StatusHandler).Methods("GET", "HEAD")
	router.HandleFunc("/_status", MethodNotAllowedHandler)
	return router
}

// StatusHandler is the basic healthcheck for the application
//
// GET /_status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	setStatusHeaders(w)

	renderer.JSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func setStatusHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "none")
}

// MethodNotAllowedHandler is an http.Handler implementation for when an HTTP method is used which isn't supported by the resource.
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	renderError(w, http.StatusMethodNotAllowed,
		"Method "+r.Method+" not allowed for <"+r.URL.RequestURI()+">")
}

func renderError(w http.ResponseWriter, status int, errorString ...string) {
	errors := newErrorInfos(errorString...)

	var message string
	if len(errors) == 1 {
		message = errorString[0]
	}
	renderer.JSON(w, status, APIResponse{
		Status:  "error",
		Message: message,
		Errors:  errors})
}

func newErrorInfos(errorString ...string) []ErrorInfo {
	errors := make([]ErrorInfo, len(errorString))
	for i, detail := range errorString {
		errors[i] = ErrorInfo{Detail: detail}
	}
	return errors
}
