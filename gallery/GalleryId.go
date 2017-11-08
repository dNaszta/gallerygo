package gallery

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"gallerygo/mongo"
)

type GalleryId struct {
	GalleryId string	`bson:"gallery_id" json:"gallery_id"`
}

func FindGalleryIds(collection *mgo.Collection) *[]GalleryId {
	var results []GalleryId
	err := collection.
		Find(bson.M{}).
		Select(bson.M{"_id" : 0, mongo.GalleryKey : 1}).
		All(&results)
	if err != nil {
		panic(err)
	}

	return &results
}