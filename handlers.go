package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func getElapsedTime(startTime time.Time) time.Duration {
	return time.Now().Sub(startTime)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "HomeHandler()")
	startTime := time.Now()

	jsonResponse, err := json.Marshal(
		struct {
			Result string
			Time   time.Duration
		}{
			Result: "Go, baby, go!",
			Time:   getElapsedTime(startTime),
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "CompareHandler()")
	startTime := time.Now()

	jsonResponse, err := json.Marshal(
		struct {
			Result bool
			Time   time.Duration
		}{
			Result: true,
			Time:   getElapsedTime(startTime),
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "ImageHandler()")
	startTime := time.Now()

	vars := mux.Vars(r)
	imageName := vars["name"]

	jsonResponse, err := json.Marshal(
		struct {
			Result string
			Time   time.Duration
		}{
			Result: imageName,
			Time:   getElapsedTime(startTime),
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}
