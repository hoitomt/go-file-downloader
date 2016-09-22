package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseUri     string                       `yaml:"base-uri"`
	ApiKey      string                       `yaml:"api-key"`
	OutputPaths map[string]map[string]string `yaml:"output-paths"`
}

func InitializeConfig(filePath string) *Config {
	config := Config{}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		panic(err)
	}

	return &config
}

func (c *Config) extractPath(environment string) string {
	environmtPaths := c.OutputPaths[environment]
	extractPath := environmtPaths["extract-path"]
	return extractPath
}
