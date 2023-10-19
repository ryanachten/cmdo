package models

import (
	"encoding/json"
	"io"
	"os"
)

type Configuration struct {
	Commands []Command `json:"commands"`
}

func ParseConfigurationFile(configPath string) (*Configuration, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = json.Unmarshal(fileContents, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
