package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Println(cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}