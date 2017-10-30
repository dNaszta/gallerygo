package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const ConfigFileName = "./config.json"

var Sizes sizesConfigs

type sizesConfigs struct {
	Sizes []sizeConfig
}

type sizeConfig struct {
	Width uint16
	Height uint16
}

func NewSizeConfig(width, height uint16) *sizeConfig {
	return &sizeConfig{
		Width: width,
		Height: height,
	}
}

func Load()  {
	file, err := ioutil.ReadFile(ConfigFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json.Unmarshal(file, &Sizes)

	if len(Sizes.Sizes) < 1 {
		fmt.Printf("Error: No image sizes config")
		os.Exit(1)
	}
}