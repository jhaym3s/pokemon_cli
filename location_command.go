package main

import (
	"fmt"
	"log"

	"github.com/jhaym3s/pokemon_cli/internal/pokeapi"
)

func LocationCallBack() error {
	pokeApiClient := pokeapi.NewClient()

	resp, err := pokeApiClient.GetLocationList()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Locations gotten")
	for _, area := range resp.Results{
		fmt.Printf("-%s \n", area)
	}

	return nil
}