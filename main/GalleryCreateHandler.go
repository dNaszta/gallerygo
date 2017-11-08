package main

import (
	"net/http"
	"gopkg.in/gorilla/mux.v1"
	"gallerygo/gallery"
	"gallerygo/rest"
)

func GalleryCreateHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	galleryId := vars["galleryId"]
	result := gallery.FindGalleryByGalleryId(GalleryCollection, galleryId)

	if result != nil {
		restError := &rest.Error{
			Message: "GalleryId is already in use",
		}
		rest.ErrorEndpoint(w, restError)
		return
	}

	gallery := &gallery.Gallery{
		GalleryId: galleryId,
	}

	gallery.Insert(GalleryCollection)
	rest.JsonEndpoint(w, gallery)
}
