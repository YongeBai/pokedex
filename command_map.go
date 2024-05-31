package main

import (
	"fmt"
)


func commandMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.GetLocations(cfg.nextURL)
	if err != nil {
		return fmt.Errorf("failed to get next locations: %w", err)
	}
	fmt.Println("Next location Areas:")
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	return nil
}