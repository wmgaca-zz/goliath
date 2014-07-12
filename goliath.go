package main

import (
	"fmt"
	"github.com/wmgaca/go-phash"
	"github.com/wmgaca/goliath/imagestore"
	"html/template"
	_ "image"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var SERVER_ADDR string = ":" + os.Getenv("PORT")

var templates = template.Must(
	template.ParseGlob(os.Getenv("GOPATH") + "/src/github.com/wmgaca/goliath/templates/*.html"))

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		templates.ExecuteTemplate(w, "error404Page", nil)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "indexPage", nil)
}

func compareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		f, _, err := r.FormFile("image")

		if err != nil {
			fmt.Fprintf(w, "ERR #1: %s", err)
			return
		}

		t, _ := ioutil.TempFile("upload", "image-")
		defer t.Close()

		_, err = io.Copy(t, f)

		if err != nil {
			fmt.Println("ERR #2:", err)
		} else {
			fmt.Println("NAME =>", t.Name())
		}

		pHash, err := phash.ImageHashDCT(t.Name())

		if err != nil {
			fmt.Println("ERR #3 =>", err)
		}

		fmt.Println("PHASH =>", pHash)

		fmt.Fprintf(w, "Hello, Go World!")
	} else {
		errorHandler(w, r, http.StatusNotFound)
	}
}

func init() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "! Provide path to the image set.")
		os.Exit(-1)
	}

	imagestorePath := os.Args[1]
	fmt.Printf("=> Init image store (path: %s)\n", imagestorePath)
	startTime := time.Now()
	imagestore.Init(imagestorePath)
	fmt.Println("   Finished in", time.Now().Sub(startTime).String())
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/compare/", compareHandler)

	fmt.Println("=> Running server on", SERVER_ADDR)

	err := http.ListenAndServe(SERVER_ADDR, nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
