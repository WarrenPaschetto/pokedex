package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no Pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cnfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...:\n", pokemon.Name)
	// logic for creating a randum number and seeing whether
	// number falls below thrshold (pokemon caught) or above (not caught)
	// based on pokemon baseexperience which means some are harder
	// to catch than others
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")
	cnfg.caughtPokemon[pokemonName] = pokemon

	return nil
}
