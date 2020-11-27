package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/GamerSenior/delivery-much/pkg/api"
)

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

func main() {
	http.HandleFunc("/recipes/", api.RecipesHandle)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
