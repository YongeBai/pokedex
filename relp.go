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
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()
		inputFields := cleanInput(input)
		if len(inputFields) == 0 {
			continue
		}
		cmd := inputFields[0]
		err := handleInput(cmd, commands, cfg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func cleanInput(input string) []string { 
	res := strings.TrimSpace(input)
	res = strings.ToLower(res)
	return strings.Fields(res)
}


func handleInput(input string, commadsMap map[string]cliCommand, cfg *config) error {
	command, ok := commadsMap[input]
	if ok {
		return command.callback(cfg)
	} else {
		return errors.New("Command not found")
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
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
	}
}

