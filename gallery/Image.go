package gallery

import (
	"gallerygo/config"
	"sync"
)

var wg sync.WaitGroup

type Image struct {
	Descriptions map[string]string	`json:"descriptions"`
	Original ImageProperty			`json:"original"`
	Instances []ImageProperty		`json:"instances"`
}

func (i *Image) CreateInstances(sizes []config.SizeConfig) {
	sizeLen := len(sizes)
	wg.Add(sizeLen)

	instanceChan := make(chan ImageProperty, sizeLen)

	for _, size := range sizes {
		go resizeToInstance(size, i.Original.Src, instanceChan)
	}

	wg.Wait()

	for pr := 0; pr < sizeLen; pr++ {
		inst := <-instanceChan
		i.Instances = append(i.Instances, inst)
	}
	close(instanceChan)
}

func resizeToInstance(size config.SizeConfig, path string, instanceCh chan ImageProperty) {
	defer wg.Done()
	inst := ImageProperty{
		Src: path,
		Width: size.Width,
		Height: size.Height,
	}
	instanceCh<-inst
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