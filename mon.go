package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func RunMongo() {
	TestGallery.Insert()
	gallery := FindGalleryByGalleryId("test_first")
	fmt.Println("Gallery:", gallery.toString())
}

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