package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Recipe struct {
	Title       string
	Href        string
	Ingredients string
	Thumbnail   string
}

type RecipeResponse struct {
	Title   string
	Version float64
	Href    string
	Results []Recipe
}

type DeliveryRecipe struct {
	Title       string
	Ingredients []string
	Link        string
	Gif         string
}

type DeliveryResponse struct {
	Keywords []string
	Recipes  []DeliveryRecipe
}

func RecipesHandle(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["i"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'i' não está presente")
		return
	}

	keywords := strings.Split(keys[0], ",")
	if len(keywords) > 3 {
		fmt.Fprintf(w, "Numero de ingredientes é maior que 3")
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://www.recipepuppy.com/api/?i=%s", keys[0]))
	if err != nil {
		fmt.Fprintf(w, "Erro ao comunicar-se com api do RecipePuppy")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Erro durante leitura da response")
	}

	var recipeResponse RecipeResponse
	err = json.Unmarshal(body, &recipeResponse)
	if err != nil {
		fmt.Printf("Erro ao fazer unmarshall do JSON\n%s", err)
	}

	dResp := DeliveryResponse{Keywords: keywords}
	for _, recipe := range recipeResponse.Results {
		ingredients := strings.Split(recipe.Ingredients, ",")
		gif, err := findGif(recipe.Title)
		if err != nil {
			fmt.Fprintf(w, "Erro ao comunicar-se com api do GIPHY")
			return
		}
		dRecipe := DeliveryRecipe{
			Title:       recipe.Title,
			Ingredients: ingredients,
			Link:        recipe.Href,
			Gif:         gif,
		}
		dResp.Recipes = append(dResp.Recipes, dRecipe)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dResp)
}

func findGif(title string) (string, error) {
	config, err := getConfiguration()
	if err != nil {
		return "", err
	}

	apiKey := config.apiKey
	params := fmt.Sprintf("?api_key=%s&q=%s&limit=1", apiKey, title)
	resp, err := http.Get("api.giphy.com/v1/gifs/search" + params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return "", nil
}

type configuration struct {
	apiKey string
}

func getConfiguration() (*configuration, error) {
	var config configuration
	file, err := os.Open("configuration/config.development.json")
	if err != nil {
		return nil, errors.New("[Error]: não foi possível abrir o arquivo de configuração")
	}
	json.NewDecoder(file).Decode(&config)
	return &config, nil
}
