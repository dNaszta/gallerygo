package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func RunMongo() {
	session, err := mgo.Dial(Configs.MongoDB.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Configs.MongoDB.Database).C(Configs.MongoDB.GalleryCollection)
	err = c.Insert(TestGallery)

	if err != nil {
		panic(err)
	}

	result := Gallery{}
	err = c.Find(bson.M{"gallery_id": "test_first"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Gallery:", result.toString())
}
