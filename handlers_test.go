package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	HomeHandler(w, req)

	if w.Code != 200 {
		log.Fatal("Response code != 200")
		t.Failed()
	}
}

func TestCompareShouldNotAllowNotPost(t *testing.T) {
	tests := []struct {
		method string
		status int
		body   string
	}{
		{"GET", http.StatusBadRequest, BadRequestMessage},
		{"PUT", http.StatusBadRequest, BadRequestMessage},
		{"POST", http.StatusOK, ""},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, "", nil)
		if err != nil {
			log.Fatal("Can't create a test request:", err)
			t.Failed()
		}

		w := httptest.NewRecorder()
		CompareHandler(w, req)

		if w.Code != test.status {
			log.Fatal("Response code != ", test.status, "(got", w.Code, "instead)")
			t.Failed()
		}

		if len(test.body) != 0 && w.Body.String() != test.body {
			log.Fatal("Body does not match.")
			t.Failed()
		}
	}
}
