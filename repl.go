package main

import (
	"bufio"
	"fmt"
	"github.com/abtsousa/gokedex/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	client pokeapi.Client
	next   *string
	prev   *string
}

var mp = map[string]cliCommand{}

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
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Printf("An error occurred: %v", err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
