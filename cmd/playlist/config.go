package playlist

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"os"
)

type Config struct {
}

var CFG Config

func readConfiguration() error {
	configFile, err := os.Open("../../configs/playlist.yaml")
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
