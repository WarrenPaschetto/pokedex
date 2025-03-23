package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cnfg *config, args ...string) error {
	if len(cnfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cnfg.caughtPokemon {
		fmt.Printf(" -%s\n", pokemon.Name)
	}

	return nil
}
