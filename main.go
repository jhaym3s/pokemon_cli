package main

import (
	"fmt"
	"log"

	"github.com/jhaym3s/pokemon_cli/internal/pokeapi"
)

func main(){

	pokeApiClient := pokeapi.NewClient()

	resp, err:= pokeApiClient.GetLocationList()

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(resp)

	//StartRepl()
}