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
	Sizes []sizeConfig
}

type sizeConfig struct {
	Width uint16
	Height uint16
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