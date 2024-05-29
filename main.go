package main

import "github.com/yongebai/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextURL *string
	prevURL *string
}

func main() {
	cfg := config {
		pokeapiClient: pokeapi.NewClient(),		
	}
	startRepl(&cfg)
}
