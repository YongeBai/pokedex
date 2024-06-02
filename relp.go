package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {		
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()

		inputFields := cleanInput(input)
		if len(inputFields) == 0 {
			continue
		}
		err := handleInput(inputFields, commands, cfg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}


func handleInput(args []string, commadsMap map[string]cliCommand, cfg *config) error {
	command, ok := commadsMap[args[0]]
	if ok {
		return command.callback(cfg, args[1:])
	} else {
		return errors.New("command not found")
	}
}

func cleanInput(input string) []string { 
	res := strings.TrimSpace(input)
	res = strings.ToLower(res)
	return strings.Fields(res)
}

type cliCommand struct {
	name string
	description string
	callback func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Display help",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the program",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "List next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "List previous 20 locations",
			callback: commandMapBack,
		},
		"explore": {
			name: "explore {location-area}",
			description: "Explore pokemon in location area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch {pokemon}",
			description: "Attempt to catch a pokemon",
			callback: commandCatch,
		},
		"inspect {pokemon}": {
			name: "inspect",
			description: "Inspect a pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "List all pokemon in pokedex",
			callback: commandPokedex,
		},
	}
}

