package pokeapi

import (
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type Client struct{
	httpClient http.Client
}
// we do not use the default http client because with that we can not set a timeout

func newClient() Client {
	return  Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}


type  PokeLocationResponse struct {
	Count    int         `json:"count"`
	Next     *string      `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}