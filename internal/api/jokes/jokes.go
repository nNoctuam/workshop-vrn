package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"workshop/internal/api"
)

const jokePath = "/api?format=json"

type JokeClient struct {
	url string
}

func NewJokeClient(url string) *JokeClient {
	return &JokeClient{
		url: url,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + jokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
