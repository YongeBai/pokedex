package main

import "fmt"

func commandExplore(cfg *config, args []string) error{
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.ExploreLocationArea(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring: %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}