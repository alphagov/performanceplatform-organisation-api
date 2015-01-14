package main

import (
	"net/http"
	"net/http/httptest"

	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {

	var testServer *httptest.Server

	BeforeEach(func() {
		testServer = testHandlerServer(newHandler())
	})

	AfterEach(func() {
		testServer.Close()
	})

	Describe("Status", func() {

		It("responds with a status of OK", func() {
			response, err := http.Get(testServer.URL + "/_status")
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			//			Expect(response).To(EqualAPIResponse(APIResponse{
			//				Status:  "ok",
			//				Message: "database seems fine"}))
		})

		It("responds to HEAD requests", func() {
			response, err := http.Head(testServer.URL + "/_status")
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("does not respond to POST requests", func() {
			response, err := http.Post(testServer.URL+"/_status",
				"application/json",
				strings.NewReader(`{"foo":"foo"}`))
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(http.StatusMethodNotAllowed))
		})

	})

})
