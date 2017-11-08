package main

import (
	"gallerygo/gallery"
	"net/http"
	"gallerygo/rest"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	results := gallery.FindGalleryIds(GalleryCollection)
	rest.JsonEndpoint(w, results)
}
