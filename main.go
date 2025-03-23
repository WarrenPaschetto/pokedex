package main

import (
	"time"

	"github.com/WarrenPaschetto/pokedexcli/internal/pokeapi"
)

func main() {
	cnfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cnfg)
}
