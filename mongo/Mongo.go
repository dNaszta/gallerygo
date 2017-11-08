package mongo

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

const GalleryKey = "gallery_id"

func GetSession(connections string) *mgo.Session {
	session, err := mgo.Dial(connections)
	if err != nil {
		panic(err)
	}

	return session
}

func GetCollection(session *mgo.Session, database, collection string) *mgo.Collection {
	session.SetMode(mgo.Monotonic, true)
	return session.DB(database).C(collection)
}

func setGalleryKeyIndex(collection *mgo.Collection) {
	index := mgo.Index{
		Key: []string{GalleryKey},
		Unique: true,
		Name: GalleryKey,
		Background: true,
	}
	err := collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func CheckAndCreateGalleryIndexes(collection *mgo.Collection) {
	indexes, err := collection.Indexes()
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
		setGalleryKeyIndex(collection)
	}
}
