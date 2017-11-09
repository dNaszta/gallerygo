package main

import (
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"net/http"
	"gopkg.in/mgo.v2"
	"os"
	"log"
	"gallerygo/mongo"
	"gallerygo/config"
)

var Configs config.Configs
var GalleryCollection *mgo.Collection

func init() {
	config.Load(config.ConfigFileName, &Configs)
	fmt.Println("Configs:", Configs.ToString())
}

func main() {
	logfile, err := os.OpenFile(
		Configs.Log,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644)
	if err != nil {
		log.Fatal(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	session := mongo.GetSession(Configs.MongoDB.ConnectionString)
	defer session.Close()
	GalleryCollection = mongo.GetCollection(
		session,
		Configs.MongoDB.Database,
		Configs.MongoDB.GalleryCollection)
	mongo.CheckAndCreateGalleryIndexes(GalleryCollection)

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).
		Methods("GET")

	r.HandleFunc("/gallery/{galleryId}", GalleryHandler).
		Methods("GET")

	r.HandleFunc("/gallery/{galleryId}/image", ImageHandler).
		Methods("POST")

	r.HandleFunc("/test_image", TestImageHandler).
		Methods("GET")

	r.HandleFunc("/gallery/{galleryId}/create", GalleryCreateHandler).
		Methods("POST")

	http.ListenAndServe(Configs.Port, r)
}