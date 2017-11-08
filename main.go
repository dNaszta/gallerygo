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
	"bufio"
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

func TestImageHandler(w http.ResponseWriter, _ *http.Request)  {
	reader, err := os.Open("./images/1280x1024.jpg")

	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	fInfo, _ := reader.Stat()
	size := fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(reader)
	fReader.Read(buf)

	// if you create a new image instead of loading from file, encode the image to buffer instead with jpg.Encode()
	// jpg.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	// Embed into an html without JPG file
	img2str := "data:image/jpg;base64," + imgBase64Str // + "\"

	w.Write([]byte(fmt.Sprintf(img2str)))
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

	r.HandleFunc("/test_image", TestImageHandler)

	http.ListenAndServe(Configs.Port, r)
}