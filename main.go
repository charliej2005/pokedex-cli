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

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return errors.New("exit failed")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

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
