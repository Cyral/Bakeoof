package main

import (
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const baseURL = "https://api.cognitive.microsoft.com/bing/v7.0/images/search"

type Response struct {
	Value []Image `json:"value"`
}

type Image struct {
	ContentURL string `json"contentUrl"`
}

// GetPictureURL returns the url for the picture
func GetPictureURL(recipe Recipe) string {
	apiKey := os.Getenv("BING_SEARCH_API_KEY")
	url := fmt.Sprintf("%s?q=%s", baseURL, url.QueryEscape(recipe.Title))

	resp, err := makeRequest(url, apiKey)
	if err != nil {
		panic(err)
	}

	json, err := parseJSON(resp)

	for _, img := range json.Value {
		if strings.HasPrefix(img.ContentURL, "https") {
			return img.ContentURL
		}
	}

	return ""
}

func parseJSON(body []byte) (*Response, error) {
	var data = &Response{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func makeRequest(url string, key string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", key)

	var client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
