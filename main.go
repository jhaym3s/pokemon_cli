package main

import "github.com/jhaym3s/pokemon_cli/internal/pokeapi"

  type Config struct{
	pokeapiClient pokeapi.Client
	nextURL *string
	prevURL *string
  }

func main(){
	cfg := Config{
		pokeapiClient:  pokeapi.NewClient(),
	}
	StartRepl(&cfg)
}