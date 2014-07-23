package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"time"
)

func getElapsedTime(startTime time.Time) time.Duration {
	return time.Now().Sub(startTime)
}

func writeJSONResponse(w http.ResponseWriter, v interface{}) {
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}

func writeErrorResponse(w http.ResponseWriter, err interface{}) {
	log.Println("ERR:", err)

	writeJSONResponse(w, struct {
		Status bool
		Error  interface{}
	}{
		Status: false,
		Error:  err,
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "HomeHandler()")
	startTime := time.Now()

	writeJSONResponse(w, struct {
		Result string
		Time   time.Duration
	}{
		Result: "Go, baby, go!",
		Time:   getElapsedTime(startTime),
	})
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "UploadImageHandler()")
	startTime := time.Now()

	zeeImage, err := NewZeeImageFromRequest(r, "image")
	if err != nil {
		log.Println("NewZeeImageFromRequest failed.")
		writeErrorResponse(w, err)
		return
	}

	go zeeImage.Compute(true)

	writeJSONResponse(w, struct {
		Status bool
		Path   string
		Time   time.Duration
	}{
		Status: true,
		Path:   zeeImage.Path,
		Time:   getElapsedTime(startTime),
	})
}

func CheckImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "CheckImageHandler()")
	startTime := time.Now()

	vars := mux.Vars(r)
	method := vars["method"]

	zeeImage, err := NewZeeImageFromRequest(r, "image")
	if err != nil {
		log.Println("NewZeeImageFromRequest failed with", err)
		writeErrorResponse(w, err)
		return
	}
	zeeImage.Compute(false)

	imageAlreadyExists := false
	switch method {
	case "phash":
		_, imageAlreadyExists = PHashMap[zeeImage.PHash]
	case "md5":
		_, imageAlreadyExists = MD5HashMap[zeeImage.MD5Hash]
	default:
		writeErrorResponse(w, "Wrong method.")
		return
	}

	writeJSONResponse(w, struct {
		Status bool
		Exists bool
		Path   string
		Time   time.Duration
	}{
		Status: true,
		Exists: imageAlreadyExists,
		Path:   zeeImage.Path,
		Time:   getElapsedTime(startTime),
	})
}
