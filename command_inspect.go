package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	cachedVal, exists := cfg.caughtPokemon[name]
	if !exists {
		fmt.Printf("No available data for %s", name)
	}
	fmt.Printf("Name: %s\n", cachedVal.Name)
	fmt.Printf("Height: %d\n", cachedVal.Height)
	fmt.Printf("Weight: %d\n", cachedVal.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range cachedVal.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pType := range cachedVal.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}
