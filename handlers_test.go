package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	HomeHandler(w, req)

	if w.Code != 200 {
		log.Fatal("Response code != 200")
		t.Failed()
	}

	if w.Body.String() != "Go, baby, go!" {
		log.Fatal("Body does not match.")
		t.Failed()

	}
}
