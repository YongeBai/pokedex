package main

import (
	"fmt"
	"log"

	"github.com/yongebai/pokedex/internal/pokeapi"
)


func commandMapBack() error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.GetLocations()
	if err != nil {
		log.Fatal(err)
	}
	if resp.Previous == nil {
		return fmt.Errorf("no previous locations")
	}
	fmt.Println("Location Areas:\n")
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	return nil
}