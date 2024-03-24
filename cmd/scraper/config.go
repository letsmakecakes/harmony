package scraper

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	URL struct {
		Hot100               string `yaml:"hot_100"`
		Billboard200         string `yaml:"billboard_200"`
		Billboard200Global   string `yaml:"billboard_200_global"`
		BillboardJapanHot100 string `yaml:"'billboard_japan_hot_100"`
	} `yaml:"url"`
}

var CFG Config

func readConfiguration() error {
	configFile, err := os.Open("../../configs/scraper.yaml")
	if err != nil {
		return fmt.Errorf("error while opening the config file: %v", err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {

		}
	}(configFile)

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&CFG)
	if err != nil {
		return fmt.Errorf("error while decoding the contents of config file: %v", err)
	}

	return nil
}

func ValidateOAuthConfig() error {
	return nil
}
