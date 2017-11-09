package config

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

type Configs struct {
	Port string
	Log string
	Sizes []SizeConfig
	MongoDB mongoSettings
}

type mongoSettings struct {
	ConnectionString string
	Database string
	GalleryCollection string
}

type SizeConfig struct {
	Suffix string
	Width uint16
	Height uint16
}

func (c *Configs) ToJSON() []byte {
	out, err := json.Marshal(c)
	if err != nil {
		panic (err)
	}
	return out
}

func (c *Configs) ToString() string {
	return string(c.ToJSON())
}

func Load(path string, configs *Configs)  {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json.Unmarshal(file, configs)

	if len(configs.Sizes) < 1 {
		fmt.Printf("Error: No image sizes config")
		os.Exit(1)
	}

	if configs.Port == "" {
		configs.Port = DefaultPort
	}
}
