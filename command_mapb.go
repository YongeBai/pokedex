package main

import (
	"fmt"
)


func commandMapBack(cfg *config, args []string) error {
	if cfg.prevURL == nil {
		return fmt.Errorf("no previous locations")
	}
	pokeapiClient := cfg.pokeapiClient	
	resp, err := pokeapiClient.GetLocations(cfg.prevURL)
	if err != nil {
		return fmt.Errorf("failed to get previous locations: %w", err)
	}
	fmt.Println("Previous Location Areas:")
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	return nil
}