package gallery

import (
	"encoding/base64"
	"strings"
	"image"
	"log"
	"os"
	"image/jpeg"
	"time"
	"fmt"
)

type ImageProperty struct {
	Src string		`json:"src"`
	Width uint16	`json:"width"`
	Height uint16	`json:"height"`
}

func Base64toJpg(data, folder, host string) (*ImageProperty, error) {
	imageProperty := &ImageProperty{}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	log.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	filename := getTimeString()
	jpgFilename := folder + "/" + filename + JPGExtension
	f, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return imageProperty, err
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		return imageProperty, err
	}

	imageProperty.Src = host + "/" + filename + JPGExtension
	imageProperty.Width = uint16(bounds.Max.X)
	imageProperty.Height = uint16(bounds.Max.Y)
	return imageProperty, err
}

func getTimeString() string {
	t := time.Now().UnixNano()
	return fmt.Sprintf("%v", t)
}