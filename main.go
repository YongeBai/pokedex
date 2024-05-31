package main

import (
	// "fmt"
	"time"

	"github.com/yongebai/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL *string
	prevURL *string
}

func main() {
	interval := time.Minute * 1
	cfg := config {
		pokeapiClient: pokeapi.NewClient(interval),		
	}
	startRepl(&cfg)
}

