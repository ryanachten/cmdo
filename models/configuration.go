package models

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"
)

type Configuration struct {
	Commands []Command `json:"commands"`
}

func GetConfigurationPath() (string, error) {
	var configPath string
	defaultValue := ""
	flag.StringVar(&configPath, "config", defaultValue, "Path for commando configuration file")
	flag.Parse()

	if configPath == defaultValue {
		return "", errors.New("set --config flag with a path to the configuration file")
	}
	return configPath, nil
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
