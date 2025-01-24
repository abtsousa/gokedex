package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, _ ...string) error {

	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not yet caught any Pokémon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pkm := range cfg.pokedex {
		fmt.Printf("- %s\n", pkm.Name)
	}

	return nil
}
