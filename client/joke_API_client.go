package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ezkripke/jokeCLI/model"
)

// JokeCategory is a type for the "category" field in jokeAPI response
type JokeCategory string

// JokeType is a type for the "Type" field in a jokeAPI response
type JokeType string

// BaseURL is the base URL of the API
const (
	BaseURL         string       = "https://sv443.net/jokeapi/v2/joke"
	DefaultCategory JokeCategory = "Any"
	DefaultType     JokeType     = "Single"
)

// JokeAPIClient is an object representing a client of the API
type JokeAPIClient struct {
	client  *http.Client
	baseURL string
}

// NewJokeAPIClient creates a JokeAPIClient
func NewJokeAPIClient() *JokeAPIClient {
	return &JokeAPIClient{
		client: &http.Client{
			Timeout: 30 * time.Second, // default
		},
		baseURL: BaseURL,
	}
}

// Fetch retrieves joke from jokeAPI based on requested category c
func (jc *JokeAPIClient) Fetch(c JokeCategory, t JokeType) (model.Output, error) {
	url := fmt.Sprintf("%s/%s?type=%s", jc.baseURL, c, t)
	fmt.Printf("Fetching joke... (%s)\n", url)
	resp, err := jc.client.Get(url)
	if err != nil {
		return model.Output{}, err
	}
	defer resp.Body.Close()

	if t == "Single" {
		var jokeResp model.Single
		if err := json.NewDecoder(resp.Body).Decode(&jokeResp); err != nil {
			return model.Output{}, err
		}
		return jokeResp.Output(), nil
	}

	// else joke format is "Twopart"

	var jokeResp model.Twopart
	if err := json.NewDecoder(resp.Body).Decode(&jokeResp); err != nil {
		return model.Output{}, err
	}
	return jokeResp.Output(), nil

}
