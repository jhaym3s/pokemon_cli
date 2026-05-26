package main

import (
	"fmt"
	"log"
)

func NextLocationCallBack(cfg *Config) error {

	resp, err := cfg.pokeapiClient.GetNextLocationList(cfg.nextURL)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Locations gotten")
	for _, area := range resp.Results {
		fmt.Printf("-%s \n", area)
	}

	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	return nil
}


func PrevLocationCallBack(cfg *Config) error {

	resp, err := cfg.pokeapiClient.GetPrevLocationList(cfg.nextURL)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Locations gotten")
	for _, area := range resp.Results {
		fmt.Printf("-%s \n", area)
	}

	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	return nil
}
