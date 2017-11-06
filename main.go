package main

import (
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"net/http"
	"gopkg.in/mgo.v2"
	"encoding/json"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	results := FindGalleryIds()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err := enc.Encode(results)
	if err != nil {
		panic(err)
	}
}

func GalleryHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	galleryId := vars["galleryId"]
	result := FindGalleryByGalleryId(galleryId)

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
	session := GetSession()
	defer session.Close()
	GalleryCollection = GetGalleryCollection(session)
	CheckAndCreateGalleryIndexes()
	// RunMongo()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/gallery/{galleryId}", GalleryHandler).Methods("GET")
	http.ListenAndServe(Configs.Port, r)
}