package main

import (
	"log"
	"net/http"

	"github.com/GamerSenior/delivery-much/pkg/api"
)

func main() {
	http.HandleFunc("/recipes/", api.RecipesHandle)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
