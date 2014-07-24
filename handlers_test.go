package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: PUT, POST
func TestHomeMethods(t *testing.T) {
	tests := []struct {
		method string
		status int
	}{
		{"GET", http.StatusOK},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, "", nil)
		if err != nil {
			log.Fatal("Can't create a test request:", err)
			t.Failed()
		}

		w := httptest.NewRecorder()
		HomeHandler(w, req)

		if w.Code != test.status {
			log.Fatal("Response code != ", test.status, "(got", w.Code, "instead)")
			t.Failed()
		}
	}
}

// TODO: PUT, POST
func TestCheckMethods(t *testing.T) {
	tests := []struct {
		method string
		status int
	}{
		{method: "POST", status: http.StatusOK},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, "", nil)
		if err != nil {
			log.Fatal("Can't create a test request:", err)
			t.Failed()
		}

		w := httptest.NewRecorder()
		CheckImageHandler(w, req)

		if w.Code != test.status {
			log.Fatal("Response code != ", test.status, "(got", w.Code, "instead)")
			t.Failed()
		}
	}
}
