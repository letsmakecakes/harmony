package oauth

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	URL    string `yaml:"url"`
	Client struct {
		ID     string `yaml:"id"`
		Secret string `yaml:"secret"`
	} `yaml:"client"`
}

var CFG Config

func readConfiguration() error {
	configFile, err := os.Open("../../configs/oauth.yaml")
	if err != nil {
		return fmt.Errorf("error while opening the config file: %v", err)
	}
	defer configFile.Close()

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
