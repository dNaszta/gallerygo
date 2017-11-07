package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const ConfigFileName = "./config.json"
const DefaultPort = ":8080"

var Configs configs

type configs struct {
	Port string
	Log string
	Sizes []sizeConfig
	MongoDB mongoSettings
}

type mongoSettings struct {
	ConnectionString string
	Database string
	GalleryCollection string
}

type sizeConfig struct {
	Width uint16
	Height uint16
}

func (c *configs) toJSON() []byte {
	out, err := json.Marshal(c)
	if err != nil {
		panic (err)
	}
	return out
}

func (c *configs) toString() string {
	return string(c.toJSON())
}

func Load()  {
	file, err := ioutil.ReadFile(ConfigFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json.Unmarshal(file, &Configs)

	if len(Configs.Sizes) < 1 {
		fmt.Printf("Error: No image sizes config")
		os.Exit(1)
	}

	if Configs.Port == "" {
		Configs.Port = DefaultPort
	}
}