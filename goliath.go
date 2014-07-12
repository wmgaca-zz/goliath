package main

import (
	"fmt"
	"github.com/wmgaca/goliath/imagestore"
	"html/template"
	"net/http"
	"os"
	"time"
)

const SERVER_ADDRESS string = ":8000"

var templates = template.Must(
	template.ParseGlob("src/github.com/wmgaca/goliath/templates/*.html"))

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
		fmt.Fprintf(w, "Hello, Go World!")
	} else {
		errorHandler(w, r, http.StatusNotFound)
	}
}

func init() {
	fmt.Println("Init:")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "  ! Provide path to the image set.")
		os.Exit(-1)
	}

	imagestorePath := os.Args[1]
	fmt.Printf("  => Imagestore (path: %s)\n", imagestorePath)
	startTime := time.Now()
	imagestore.Init(imagestorePath)
	fmt.Println("     Finished in", time.Now().Sub(startTime).String())
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/compare/", compareHandler)

	fmt.Println("Listening on", SERVER_ADDRESS)
	http.ListenAndServe(SERVER_ADDRESS, nil)
}
