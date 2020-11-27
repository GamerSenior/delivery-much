package internal

import (
	"encoding/json"
	"errors"
	"os"
)

// configuration é um struct que representa o json de configuração
type Configuration struct {
	apiKey string
}

func GetConfiguration() (*Configuration, error) {
	var config Configuration
	file, err := os.Open("configuration/config.development.json")
	if err != nil {
		return nil, errors.New("[Error]: não foi possível abrir o arquivo de configuração")
	}
	json.NewDecoder(file).Decode(&config)
	return &config, nil
}
