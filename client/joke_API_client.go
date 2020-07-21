package client

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ezkripke/jokeCLI/model"
)

// JokeCategory is a type for the "category" field in jokeAPI response
type JokeCategory string

// BaseURL is the base URL of the API
const (
	BaseURL         string       = "https://sv443.net/jokeapi/v2/joke"
	DefaultCategory JokeCategory = "Programming,Miscellaneous"
	DefaultBTags    string       = "nsfw,religious,political,racist,sexist"
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
func (jc *JokeAPIClient) Fetch(c JokeCategory) (model.Output, error) {
	// choose randomly between "single" and "twopart" joke types so that we know
	// which struct from `model` to read API response into
	ts := []string{"single", "twopart"}
	t := ts[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(ts))]

	url := fmt.Sprintf(
		"%s/%s?type=%s?blacklistFlags=%s",
		jc.baseURL, c, t, DefaultBTags,
	)
	fmt.Printf("Fetching joke... \n(%s)\n\n", url)

	resp, err := jc.client.Get(url)
	if err != nil {
		return model.Output{}, err
	}
	defer resp.Body.Close()

	if t == "single" {
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
