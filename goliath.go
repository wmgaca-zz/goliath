package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var DisapprovalLook string = "ಠ_ಠ"

var ServerAddr string

func ExitWithErr(message string) {
	log.Println(DisapprovalLook, message)
	os.Exit(-1)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "HomeHandler()")
	fmt.Fprintf(w, "Go, baby, go!")
}

func init() {
	ServerAddr = ":" + os.Getenv("PORT")
	if len(ServerAddr) == 1 {
		ExitWithErr("Server port not set.")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	http.Handle("/", router)

	log.Println("Runnig server on", ServerAddr)
	http.ListenAndServe(ServerAddr, nil)
}
