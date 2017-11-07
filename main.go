package main

import (
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"net/http"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"os"
	"log"
	"encoding/base64"
	"strings"
	"image"
	_ "image/jpeg"
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
	var src SourceImage

	for key, _ := range r.Form {
		err := json.Unmarshal([]byte(key), &src)
		if err != nil {
			panic(err)
		}
	}

	b64data := src.Source[strings.IndexByte(src.Source, ',')+1:]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64data))
	img, _, err := image.Decode(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Before bound")
	bound := img.Bounds()
	log.Println(bound)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonEndpoint(w, src)
}

func itemNotFoundEndpoint(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func jsonEndpoint(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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