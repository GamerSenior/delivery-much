package config

import (
	"encoding/json"
	"os"
)

// configuration é um struct que representa o json de configuração
type Configuration struct {
	ApiKey string
}

func GetConfiguration() (*Configuration, error) {
	var config Configuration
	file, err := os.Open("config/config.development.json")
	if err != nil {
		return nil, err
	}
	json.NewDecoder(file).Decode(&config)
	return &config, nil
}
