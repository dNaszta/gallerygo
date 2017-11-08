package main

import (
	"encoding/json"
	"gallerygo/gallery"
	"strings"
	"net/http"
	"gallerygo/rest"
)

func ImageHandler(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var src gallery.SourceImage
	err := decoder.Decode(&src)
	if err != nil {
		panic(err)
	}

	b64data := src.Source[strings.IndexByte(src.Source, ',')+1:]
	imageProperty, err := gallery.Base64toJpg(
		b64data,
		Configs.ImageFolder,
		Configs.ImageHost)
	if err != nil {
		panic(err)
	}

	rest.JsonEndpoint(w, imageProperty)
}
