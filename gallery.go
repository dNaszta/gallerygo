package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Gallery struct {
	GalleryId string	`bson:"gallery_id" json:"gallery_id"`
	Images []Image		`json:"images"`
}

type Image struct {
	Descriptions map[string]string	`json:"descriptions"`
	Original ImageProperty			`json:"original"`
	Instances []ImageProperty		`json:"instances"`
}

type ImageProperty struct {
	Src string		`json:"gallery_id"`
	Width uint16	`json:"width"`
	Height uint16	`json:"height"`
}

func (g *Gallery) toJSON() []byte {
	out, err := json.Marshal(g)
	if err != nil {
		panic (err)
	}
	return out
}

func (g *Gallery) toString() string {
	return string(g.toJSON())
}

func (g *Gallery) Insert() {
	session, err := mgo.Dial(Configs.MongoDB.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(Configs.MongoDB.Database).C(Configs.MongoDB.GalleryCollection)
	err = collection.Insert(g)

	if err != nil {
		panic(err)
	}
}

func FindGalleryByGalleryId(gallery_id string) *Gallery {
	session, err := mgo.Dial(Configs.MongoDB.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(Configs.MongoDB.Database).C(Configs.MongoDB.GalleryCollection)

	result := Gallery{}
	err = collection.Find(bson.M{"gallery_id": "test_first"}).One(&result)
	if err != nil {
		panic(err)
	}

	return &result
}

var TestGallery = Gallery {
	GalleryId: "test_first",
	Images: []Image {
		Image {
			Descriptions: map[string]string {
				"hu":"Hungarian Description",
				"en":"English Description",
			},
			Original: ImageProperty {
				Src: "https://www.placecage.com/200/300",
				Width: 200,
				Height: 300,
			},
			Instances: []ImageProperty{
				ImageProperty {
					Src: "https://www.placecage.com/640/480",
					Width: 640,
					Height: 480,
				},
				ImageProperty {
					Src: "https://www.placecage.com/1024/768",
					Width: 1024,
					Height: 768,
				},
			},
		},
	},
}
