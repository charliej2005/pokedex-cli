package main

import (
	"time"

	pokeapi "github.com/charliej2005/pokedex-cli/internal"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	StartRepl(cfg)
}
