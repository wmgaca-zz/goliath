package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var serverAddr string

func init() {
	serverAddr = ":" + os.Getenv("PORT")
	if len(serverAddr) == 1 {
		ExitWithErr("Server port not set.")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/compare/", CompareHandler).Methods("POST")
	router.HandleFunc("/image/{name}/", ImageHandler).Methods("GET")
	http.Handle("/", router)

	log.Println("Runnig server on", serverAddr)
	http.ListenAndServe(serverAddr, nil)
}
