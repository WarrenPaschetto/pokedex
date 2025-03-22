package main

import (
	"errors"
	"fmt"
)

func commandExplore(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	locationArea, err := cnfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
