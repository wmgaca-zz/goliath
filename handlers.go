package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r, "HomeHandler()")
	fmt.Fprintf(w, "Go, baby, go!")
}
