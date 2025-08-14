package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return errors.New("exit failed")
}

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func main() {
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
			userCommand = strings.ToLower(cleanedInput[0])
		}
		command, ok := commands[userCommand]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("error executing command: %v", err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
