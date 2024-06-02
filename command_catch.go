package main

import (
	"fmt"
	"math/rand"

	"github.com/yongebai/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 { 
		return fmt.Errorf("missing pokemon name")
	}
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing Pokeball at %s...\n", args[0])
	if attemptCatch(resp.BaseExperience) {
		fmt.Printf("Caught %s!\n", args[0])
		cfg.pokedex[args[0]] = Pokemon{
			Name: args[0],
			Height: resp.Height,
			Weight: resp.Weight,
			Stats: resp.Stats,
			Types: resp.Types,
			Moves: getRandomMoves(resp),
		}
	} else {
		fmt.Printf("%s broke free!\n", args[0])	
	}
	return nil
}

func attemptCatch(baseExperience int) bool {
	chanceAtCatch := 800 - baseExperience
	randomNumber := rand.Intn(1000)
	return randomNumber < chanceAtCatch
}

func getRandomMoves(pokemonResponse pokeapi.PokemonResponse) []string {	
	var randomMoveIdx int	
	moves := make([]string, 4)

	for i := 0; i < 4; i++ {
		randomMoveIdx = rand.Intn(len(pokemonResponse.Moves))
		moves[i] = pokemonResponse.Moves[randomMoveIdx].Move.Name
	}
	return moves
}