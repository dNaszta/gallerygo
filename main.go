package main

import (
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"net/http"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"os"
	"log"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	results := FindGalleryIds()

	jsonEndpoint(w, results)
}

func GalleryHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	galleryId := vars["galleryId"]
	result := FindGalleryByGalleryId(galleryId)
	if result == nil {
		itemNotFoundEndpoint(w)
	} else {
		jsonEndpoint(w, result)
	}
}

func ImageHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	// log.Println(r.Form)
	var src SourceImage

	for key, _ := range r.Form {
		err := json.Unmarshal([]byte(key), &src)
		if err != nil {
			panic(err)
		}
	}

	jsonEndpoint(w, src)
}

func itemNotFoundEndpoint(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
}

func jsonEndpoint(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err := enc.Encode(result)
	if err != nil {
		panic(err)
	}
}

func init() {
	Load()
	fmt.Println("Configs:", Configs.toString())
}

var GalleryCollection *mgo.Collection

func main() {
	logfile, err := os.OpenFile(Configs.Log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	session := GetSession()
	defer session.Close()
	GalleryCollection = GetGalleryCollection(session)
	CheckAndCreateGalleryIndexes()
	// RunMongo()

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).
		Methods("GET")

	r.HandleFunc("/gallery/{galleryId}", GalleryHandler).
		Methods("GET")

	r.HandleFunc("/gallery/{galleryId}/image", ImageHandler).
		Methods("POST")

	http.ListenAndServe(Configs.Port, r)
}