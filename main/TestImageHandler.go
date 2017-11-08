package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"net/http"
	"encoding/base64"
	"gallerygo/gallery"
)

func TestImageHandler(w http.ResponseWriter, _ *http.Request)  {
	reader, err := os.Open("./images/1280x1024.jpg")

	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	fInfo, _ := reader.Stat()
	size := fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(reader)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	img2str := gallery.Base64JpgStart + imgBase64Str

	w.Write([]byte(fmt.Sprintf(img2str)))
}
