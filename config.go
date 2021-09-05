package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

func readConfig() (*Config, error) {
	var config *Config

	configFile, err := os.Open("config.json")
	if err != nil {
		color.Red("[ERROR] config.readConfig Error Opening config.json: %s", err)
		return nil, err
	}
	defer configFile.Close()

	configData, _ := ioutil.ReadAll(configFile)

	json.Unmarshal(configData, &config)

	return config, nil
}
