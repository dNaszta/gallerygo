package gallery

import (
	"gallerygo/config"
	"sync"
	"os"
	"log"
	"image/jpeg"
	"github.com/nfnt/resize"
	"strings"
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

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(uint(size.Width), uint(size.Height), img, resize.Lanczos3)

	baseFilename := strings.TrimLeft(path, config.Folder)
	jpgFilename := config.Folder + size.Suffix + "/" + size.Suffix + baseFilename
	out, err := os.Create(jpgFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	inst := ImageProperty{
		Src: jpgFilename,
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