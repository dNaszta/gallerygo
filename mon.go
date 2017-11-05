package main

import (
	"fmt"
)

func RunMongo() {
	TestGallery.Insert()
	gallery := FindGalleryByGalleryId("test_first")
	fmt.Println("Gallery:", gallery.toString())
}
