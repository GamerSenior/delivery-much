package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/GamerSenior/delivery-much/internal/config"
)

type baseResponse struct {
	Data []gif `json:"data"`
}

type gif struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	BitlyURL string `json:"bitly_url"`
}

// GetGifURLByTitle recebe uma string title e retorna a url
// do primeiro GIF encontrado no GIPHY
func GetGifURLByTitle(title string) (string, error) {
	config, err := config.GetConfiguration()
	if err != nil {
		return "", err
	}

	apiKey := config.ApiKey
	baseURL := "http://api.giphy.com/v1/gifs/search?"
	v := url.Values{}
	v.Set("api_key", apiKey)
	v.Add("q", title)
	v.Add("limit", "1")
	baseURL += v.Encode()
	resp, err := http.Get(baseURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var bResp baseResponse
	if err := json.Unmarshal(data, &bResp); err != nil {
		return "", err
	}

	if len(bResp.Data) > 0 {
		return bResp.Data[0].URL, nil
	}
	return "", nil
}
