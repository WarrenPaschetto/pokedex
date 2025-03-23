package main

import (
	"errors"
	"fmt"
)

func commandInspect(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no Pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, exists := cnfg.caughtPokemon[pokemonName]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}

	return nil
}
