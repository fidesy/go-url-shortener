package utils

import (
	"os"

	"github.com/fidesy/go-url-shortener/internal/config"
	"gopkg.in/yaml.v2"
)

func LoadConfig(filepath string) (*config.Config, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var conf *config.Config
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
