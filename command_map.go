package main

import (
	"fmt"
	"log"
)


func commandMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.GetLocations(cfg.nextURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Next ation Areas:\n")
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	return nil
}