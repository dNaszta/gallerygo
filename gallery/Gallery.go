package gallery

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

const JPGExtension = ".jpg"
const Base64JpgStart = "data:image/jpg;base64,"

type Gallery struct {
	GalleryId string	`bson:"gallery_id" json:"gallery_id"`
	Images []Image		`json:"images"`
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

func (g *Gallery) Insert(collection *mgo.Collection) {
	err := collection.Insert(g)

	if err != nil {
		panic(err)
	}
}

func FindGalleryByGalleryId(collection *mgo.Collection, galleryId string) *Gallery {
	gallery := Gallery{}
	err := collection.
		Find(bson.M{"gallery_id" : galleryId}).
		One(&gallery)
	if err != nil {
		return nil
	}

	return &gallery
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
