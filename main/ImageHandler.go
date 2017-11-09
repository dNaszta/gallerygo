package main

import (
	"encoding/json"
	"gallerygo/gallery"
	"strings"
	"net/http"
	"gallerygo/rest"
	"gopkg.in/gorilla/mux.v1"
)

func ImageHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	galleryId := vars["galleryId"]
	galleryItem := gallery.FindGalleryByGalleryId(GalleryCollection, galleryId)

	if galleryItem == nil {
		restError := &rest.Error{
			Message: "Invalid gallery_id",
		}
		rest.ErrorEndpoint(w, restError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var src gallery.SourceImage
	err := decoder.Decode(&src)
	if err != nil {
		panic(err)
	}

	b64data := src.Source[strings.IndexByte(src.Source, ',')+1:]
	imageProperty, err := gallery.Base64toJpg(b64data)
	if err != nil {
		panic(err)
	}

	image := gallery.CreateImageByPropertyAndSource(imageProperty, src)
	image.CreateInstances(Configs.Sizes)
	galleryItem.AddNewOriginalImage(image)
	galleryItem.Update(GalleryCollection)

	rest.JsonEndpoint(w, imageProperty)
}
