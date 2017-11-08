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
	"bufio"
	"image/jpeg"
	"time"
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
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var src SourceImage
	err := decoder.Decode(&src)
	if err != nil {
		panic(err)
	}

	b64data := src.Source[strings.IndexByte(src.Source, ',')+1:]
	imageProperty, err := base64toJpg(b64data)
	if err != nil {
		panic(err)
	}

	jsonEndpoint(w, imageProperty)
}

func getTimeString() string {
	t := time.Now().UnixNano()
	return fmt.Sprintf("%v", t)
}

func base64toJpg(data string) (*ImageProperty, error) {
	imageProperty := &ImageProperty{}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	log.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	filename := getTimeString()
	jpgFilename := Configs.ImageFolder + "/" + filename + JPGExtension
	f, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return imageProperty, err
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		return imageProperty, err
	}

	imageProperty.Src = Configs.ImageHost + "/" + filename + JPGExtension
	imageProperty.Width = uint16(bounds.Max.X)
	imageProperty.Height = uint16(bounds.Max.Y)
	return imageProperty, err
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

	fReader := bufio.NewReader(reader)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	img2str := Base64JpgStart + imgBase64Str

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
	logfile, err := os.OpenFile(
		Configs.Log,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644)
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