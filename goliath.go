package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var ServerAddr string

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "HomeHandler()")
	fmt.Fprintf(w, "Go, baby, go!")
}

func init() {
	ServerAddr = ":" + os.Getenv("PORT")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	http.Handle("/", router)
	http.ListenAndServe(ServerAddr, nil)
}
