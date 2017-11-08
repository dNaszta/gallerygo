package main

import (
	"gopkg.in/gorilla/mux.v1"
	"net/http"
	"gallerygo/rest"
	"gallerygo/gallery"
)

func GalleryHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	galleryId := vars["galleryId"]
	result := gallery.FindGalleryByGalleryId(GalleryCollection, galleryId)
	if result == nil {
		rest.ItemNotFoundEndpoint(w)
	} else {
		rest.JsonEndpoint(w, result)
	}
}
