package gallery

import (
	"gallerygo/config"
	"log"
)

type Image struct {
	Descriptions map[string]string	`json:"descriptions"`
	Original ImageProperty			`json:"original"`
	Instances []ImageProperty		`json:"instances"`
}

func (i *Image) CreateInstances(sizes []config.SizeConfig) {
	log.Println("Sizes:", len(sizes))
	for _, size := range sizes {
		instance := resizeToInstance(size, i.Original.Src)
		i.Instances = append(i.Instances, instance)
	}
}

func resizeToInstance(size config.SizeConfig, path string) ImageProperty {
	return ImageProperty{
		Src: path,
		Width: size.Width,
		Height: size.Height,
	}
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