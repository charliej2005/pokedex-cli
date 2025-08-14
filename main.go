package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		var command string
		if len(cleanedInput) == 0 {
			command = ""
		} else {
			command = strings.ToLower(cleanedInput[0])
		}
		fmt.Printf("Your command was: %v\n", command)
	}
}
