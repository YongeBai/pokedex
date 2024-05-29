package main

import (
	"fmt"
	"log"
)


func commandMapBack(cfg *config) error {
	if cfg.prevURL == nil {
		return fmt.Errorf("no previous locations")
	}
	pokeapiClient := cfg.pokeapiClient	
	resp, err := pokeapiClient.GetLocations(cfg.prevURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Previous Location Areas:\n")
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	return nil
}