package main

import "fmt"

func commandExplore(cfg *config, arg string) error {
	locationsRes, err := cfg.pokeapiClient.ListFoundPokemons(cfg.nextLocationsURL, arg)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Found Pokemon:")

	for _, enc := range locationsRes.PokemonEncounters {
		fmt.Println(" - " + enc.Pokemon.Name)
	}

	return nil
}
