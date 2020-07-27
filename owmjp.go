package owmjp

import (
	"net/http"
)

const (
	BaseURL = "http://api.openweathermap.org/data/2.5"
)

type API struct {
	APIKey     string
	HTTPClient *http.Client
	baseURL    string
}

var (
	apiKey string
)

func NewAPI() *API {
	c := &API{
		APIKey:     apiKey,
		HTTPClient: http.DefaultClient,
		baseURL:    BaseURL,
	}
	return c
}

func SetAPIKey(_apiKey string) {
	apiKey = _apiKey
}

func NewAPIWithKey(_apiKey string) *API {
	c := NewAPI()
	c.APIKey = _apiKey
	return c
}
