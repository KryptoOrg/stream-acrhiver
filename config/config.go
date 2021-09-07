package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Address   string   `yaml:"address"`
	Filename  string   `yaml:"filename"`
	Directory string   `yaml:"directory"`
	Frequency int64    `yaml:"frequency_seconds"`
	Symbols   []string `yaml:"symbols"`
}

func NewConfig(configFilename string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(configFilename)
	if err != nil {
		log.Error("Error while reading config file: ", err)
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Error("Unmarshal: %v", err)
		return nil, err
	}

	return &config, nil
}
