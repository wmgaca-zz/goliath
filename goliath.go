package main

import (
	"github.com/gorilla/mux"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
	"net/http"
	"os"
)

const StaticDir = "./static/"

const S3ImageBucketName = "goliath-images"

var S3Connection *s3.S3

var S3GoliathImagesBucket *s3.Bucket

var ServerAddr string

func init() {
	// Server port
	ServerAddr = ":" + os.Getenv("PORT")
	if len(ServerAddr) == 1 {
		ExitWithErr("Server port not set.")
	}

	// Amazon S3
	auth, err := aws.EnvAuth()
	if err != nil {
		ExitWithErr("AWS, can't auth")
	}
	euwest := aws.USEast
	S3Connection = s3.New(auth, euwest)
	S3GoliathImagesBucket = S3Connection.Bucket(S3ImageBucketName)
}

func configureRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/compare/", CompareHandler).Methods("POST")
	router.HandleFunc("/image/{name}/", ImageHandler).Methods("GET")
	router.HandleFunc("/list/", ListBucketHandler).Methods("GET")

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
