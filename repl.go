package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zsoltdzsugan/pokedexcli/internal/pokeapi"
)

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startREPL(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		command, exists := cfg.commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show found pokemon in that area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback:    commandInspect,
		},
	}
}
