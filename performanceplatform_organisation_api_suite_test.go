package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
	"testing"

)

func TestPerformanceplatformOrganisationApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PerformanceplatformOrganisationApi Suite")
}


func testHandlerServer(handler interface{}) *httptest.Server {
	var h http.Handler
	switch handler := handler.(type) {
	case http.Handler:
		h = handler
	case func(http.ResponseWriter, *http.Request):
		h = http.HandlerFunc(handler)
	default:
		// error
		panic("handler cannot be used in an HTTP Server")
	}
	return httptest.NewServer(h)
}
