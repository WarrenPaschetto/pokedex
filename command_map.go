package main

import (
	"errors"
	"fmt"
)

func commandMapf(cnfg *config, args ...string) error {
	locationsResp, err := cnfg.pokeapiClient.ListLocations(cnfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cnfg.nextLocationsURL = locationsResp.Next
	cnfg.prevLocationsURL = locationsResp.Previous
	fmt.Println("Location areas:")
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cnfg *config, args ...string) error {
	if cnfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cnfg.pokeapiClient.ListLocations(cnfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cnfg.nextLocationsURL = locationsResp.Next
	cnfg.prevLocationsURL = locationsResp.Previous
	fmt.Println("Location areas:")
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
