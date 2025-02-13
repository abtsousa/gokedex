package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abtsousa/gokedex/internal/cache"
	"github.com/abtsousa/gokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	client  pokeapi.Client
	cache   *cache.Cache
	pokedex pokedex
	next    *string
	prev    *string
}

var mp map[string]cliCommand

func startRepl(cfg *Config) {
	mp = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show next 20 locations",
			callback:    commandMapN,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 locations",
			callback:    commandMapP,
		},
		"explore": {
			name:        "explore",
			description: "Show all the Pokémon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokémon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokémon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		in.Scan()
		in := strings.Fields(strings.ToLower(in.Text()))

		cmd, ok := mp[in[0]]
		if !ok {
			fmt.Printf("Unknown command %s\n", in[0])
		} else {
			err := cmd.callback(cfg, in[1:]...)
			if err != nil {
				fmt.Printf("An error occurred: %v\n", err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
