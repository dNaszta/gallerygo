package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

const GalleryKey = "gallery_id"

func GetSession() *mgo.Session {
	session, err := mgo.Dial(Configs.MongoDB.ConnectionString)
	if err != nil {
		panic(err)
	}

	return session
}

func GetGalleryCollection(session *mgo.Session) *mgo.Collection {
	session.SetMode(mgo.Monotonic, true)
	return session.DB(Configs.MongoDB.Database).C(Configs.MongoDB.GalleryCollection)
}

func setGalleryKeyIndex() {
	index := mgo.Index{
		Key: []string{GalleryKey},
		Unique: true,
		Name: GalleryKey,
		Background: true,
	}
	err := GalleryCollection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func CheckAndCreateGalleryIndexes() {
	indexes, err := GalleryCollection.Indexes()
	if err != nil {
		panic(err)
	}

	isGalleryIdIndexed := false
	for _, index := range indexes {
		fmt.Println(index.Name)
		if index.Name == "gallery_id" {
			isGalleryIdIndexed = true
		}
	}

	fmt.Println(isGalleryIdIndexed)

	if isGalleryIdIndexed == false {
		// setGalleryKeyIndex()
	}
}

func RunMongo() {
	TestGallery.Insert()
	gallery := FindGalleryByGalleryId("test_first")
	fmt.Println("Gallery:", gallery.toString())
}