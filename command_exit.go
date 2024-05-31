package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args []string) error {
	fmt.Println()
	fmt.Println("Exiting the Pokedex")
	os.Exit(0)
	return nil
}
