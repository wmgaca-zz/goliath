package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const StaticDir = "./static/"

var ServerAddr string

func init() {
	// Static file
	_, err := os.Stat(StaticDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(StaticDir, 0644)
		if err != nil {
			ExitWithErr("Can't create static dir")
		}
	}

	// Server port
	ServerAddr = ":" + os.Getenv("PORT")
	if len(ServerAddr) == 1 {
		ExitWithErr("Server port not set.")
	}
}

func configureRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/compare/", CompareHandler).Methods("POST")
	router.HandleFunc("/image/{name}/", ImageHandler).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir(StaticDir))))

	return router
}

func main() {
	http.Handle("/", configureRouter())

	log.Println("Runnig server on", ServerAddr)
	http.ListenAndServe(ServerAddr, nil)
}
