package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
	Port  string `yaml:"port"`
}

const (
	filename = "../../config.sample.yml"
)

func InitConfig() *Config {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading Config file: %s", err)
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML file: %s", err)
	}
	return &config
}
