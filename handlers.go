package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const BadRequestMessage string = "These are not the droids you are looking for."

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, BadRequestMessage)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "HomeHandler()")
	fmt.Fprintf(w, "Go, baby, go!")
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "CompareHandler()")
	startTime := time.Now()

	if r.Method == "POST" {
		fmt.Fprintf(w, "YES %s", time.Now().Sub(startTime).String())
	} else {
		badRequest(w, r)
	}
}
