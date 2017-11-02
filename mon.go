package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func RunMongo() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("gallery")
	err = c.Insert(&Gallery{"test_1"},&Gallery{"test_2"})

	if err != nil {
		panic(err)
	}

	result := Gallery{}
	err = c.Find(bson.M{"gallery_id": "test_1"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Gallery:", result)
}