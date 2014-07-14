package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getElapsedTime(startTime time.Time) time.Duration {
	return time.Now().Sub(startTime)
}

func WriteJSONResponse(w http.ResponseWriter, v interface{}) {
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "HomeHandler()")
	startTime := time.Now()

	WriteJSONResponse(w, struct {
		Result string
		Time   time.Duration
	}{
		Result: "Go, baby, go!",
		Time:   getElapsedTime(startTime),
	})
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "CompareHandler()")
	startTime := time.Now()

	imageFile, imageFileHeader, err := r.FormFile("image")
	if err != nil {
		WriteJSONResponse(w, struct{ error error }{err})
		return
	}

	tempFile, err := ioutil.TempFile(StaticDir, imageFileHeader.Filename+"-")
	defer tempFile.Close()
	if err != nil {
		WriteJSONResponse(w, struct{ error error }{err})
		return
	}

	_, err = io.Copy(tempFile, imageFile)
	if err != nil {
		WriteJSONResponse(w, struct{ error error }{err})
		return
	}

	WriteJSONResponse(w, struct {
		FilePath string
		Time     time.Duration
	}{
		FilePath: tempFile.Name(),
		Time:     getElapsedTime(startTime),
	})
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "ImageHandler()")
	startTime := time.Now()

	vars := mux.Vars(r)
	imageName := vars["name"]

	WriteJSONResponse(w, struct {
		Result string
		Time   time.Duration
	}{
		Result: imageName,
		Time:   getElapsedTime(startTime),
	})
}
