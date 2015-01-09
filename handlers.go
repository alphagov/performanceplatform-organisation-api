package main

import (
	"net/http"

	"github.com/go-martini/martini"
)

func newHandler() http.Handler {
	return martini.Classic()
}
