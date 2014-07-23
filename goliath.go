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

var Debug bool = false

func init() {
	// Debug mode
	if len(os.Getenv("DEBUG")) > 0 {
		log.Println("Debug mode on.")
		Debug = true
	}

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
	router.HandleFunc("/1/upload/", UploadImageHandler).Methods("POST")
	router.HandleFunc("/1/check/{method:[a-z0-9]+}/", CheckImageHandler).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir(StaticDir))))

	return router
}

func main() {
	http.Handle("/", configureRouter())

	log.Println("Runnig server on", ServerAddr)

	if Debug {
		panic(http.ListenAndServe(ServerAddr, nil))
	} else {
		http.ListenAndServe(ServerAddr, nil)
	}
}
