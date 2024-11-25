package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/scottstensland/uitwaaien/models"
	"gopkg.in/yaml.v3"
)

func ReadConfig() (models.Constants, error) {
	fmt.Println("Reading config")

	// get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return models.Constants{}, err
	}

	fmt.Printf("\nOK here is current working directory: %s\n\n", dir)

	// get the absolute path to the config file
	configFile := filepath.Join(dir, "configs", "config.yaml")

	// read the file
	data, err := os.ReadFile(configFile)
	if err != nil {
		return models.Constants{}, err
	}

	fmt.Printf("\nOK here is input config file: %s\n\n", configFile)

	// var config models.Constants
	var cfg struct {
		Constants models.Constants `yaml:"constants"`
	}

	// unmarshal the yaml file into the config struct
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return models.Constants{}, err
	}
	fmt.Printf("here is config: %v\n\n", cfg)

	return cfg.Constants, nil
}
