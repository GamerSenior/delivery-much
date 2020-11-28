package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// DeliveryRecipe estrutura de retorno da API
// Contem os dados da receita
type DeliveryRecipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	GIF         string   `json:"gif"`
}

// DeliveryResponse estrutura base de retorno da API
// Contem palavras chafes utilizadas na busca bem como
// as receitas encontradas
type DeliveryResponse struct {
	Keywords []string
	Recipes  []DeliveryRecipe
}

// RecipesHandle - handler function responsável por retornar as receitas
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
		gif, err := GetGifURLByTitle(recipe.Title)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Erro ao comunicar-se com api do GIPHY")
			return
		}
		dRecipe := DeliveryRecipe{
			Title:       recipe.Title,
			Ingredients: ingredients,
			Link:        recipe.Href,
			GIF:         gif,
		}
		dResp.Recipes = append(dResp.Recipes, dRecipe)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dResp)
}
