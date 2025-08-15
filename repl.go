package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/charliej2005/pokedex-cli/internal"
)

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		var userCommand string
		if len(cleanedInput) == 0 {
			userCommand = ""
		} else {
			userCommand = cleanedInput[0]
		}
		command, ok := commands[userCommand]
		if ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Printf("error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func getCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
}
