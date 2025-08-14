package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return errors.New("exit failed")
}
