package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf(" - %s\n", typ.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
