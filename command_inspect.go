package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		return fmt.Errorf("pokemon %s not found in pokedex", args[0])
	}
	fmt.Printf("Name: %s\n", args[0])
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}