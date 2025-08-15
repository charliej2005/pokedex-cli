package main

import (
	"time"

	"github.com/charliej2005/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	StartRepl(cfg)
}
