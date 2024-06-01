package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfs *config, args []string) error {
	pokeapiClient := cfs.pokeapiClient
	resp, err := pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing Pokeball at %s...\n", args[0])
	if attemptCatch(args[0], resp.BaseExperience) {
		fmt.Printf("Caught %s!\n", args[0])
		cfs.pokedex[args[0]] = Pokemon{
			Name: args[0],
			Height: resp.Height,
			Weight: resp.Weight,
			Stats: resp.Stats,
			Types: resp.Types,
		}
	} else {
		fmt.Printf("%s broke free!\n", args[0])	
	}
	return nil
}

func attemptCatch(pokemonName string, baseExperience int) bool {
	chanceAtCatch := 1000 - baseExperience
	randomNumber := rand.Intn(1000)
	return randomNumber < chanceAtCatch
}