package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseWithData struct {
	Data []CatBreed `json:"data"`
}

type CatBreed struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

func GetBreedsFromURL(url string) (ResponseWithData, error) {
	res, err := http.Get(url)
	if err != nil {
		return ResponseWithData{}, fmt.Errorf("api call error: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseWithData{}, fmt.Errorf("read body error: %v", err)
	}

	var catBreed ResponseWithData
	err = json.Unmarshal(body, &catBreed)
	if err != nil {
		return ResponseWithData{}, fmt.Errorf("json unmarshal error: %v", err)
	}

	return catBreed, nil
}
