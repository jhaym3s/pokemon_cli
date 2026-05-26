package pokeapi

import (
	"net/http"
	"time"

	"github.com/jhaym3s/pokemon_cli/cache"
)


const baseUrl = "https://pokeapi.co/api/v2/"


type Client struct{
	httpClient http.Client
	cache cache.Cache
}
// we do not use the default http client because with that we can not set a timeout

func NewClient(interval time.Duration) Client {
	return  Client{
		cache: *cache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
