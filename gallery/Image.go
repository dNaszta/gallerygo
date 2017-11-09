package gallery

import (
	"gallerygo/config"
)

type Image struct {
	Descriptions map[string]string	`json:"descriptions"`
	Original ImageProperty			`json:"original"`
	Instances []ImageProperty		`json:"instances"`
}

func (i *Image) CreateInstances(sizes []config.SizeConfig) {

}

func CreateImageByPropertyAndSource(property *ImageProperty, source SourceImage) (Image){
	image := Image{
		Descriptions: source.Descriptions,
		Original: ImageProperty{
			Src: property.Src,
			Width: property.Width,
			Height: property.Height,
		},
	}
	return image
}