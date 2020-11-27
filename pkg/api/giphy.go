package api

import (
	"fmt"
	"net/http"
)

func GetGifURLByTitle(title string) (string, error) {
	config, err := GetConfiguration()
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
