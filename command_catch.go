package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func isCaught(baseExp int) bool {
	const K = 0.015
	const BASE_CHANCE = 1.
	// logistic function
	return rand.Float64() <= BASE_CHANCE*(1-1/(1+math.Exp(-K*float64(baseExp-160))))
}

func commandCatch(cfg *Config, params ...string) error {

	if len(params) == 0 {
		return fmt.Errorf("You must provide a valid Pokémon name.")
	}

	pkm, err := cfg.client.ListPokemonDetails(params[0], cfg.cache)
	if err != nil {
		return err
	}

	if len(pkm.Name) == 0 {
		fmt.Println("No Pokémon found!")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])
	if isCaught(pkm.BaseExperience) {
		fmt.Printf("%s was caught!\n", params[0])
		if _, bool := cfg.pokedex[pkm.Name]; !bool {
			fmt.Println("You may now inspect it with the inspect command.")
			cfg.pokedex[pkm.Name] = pkm
		}
	} else {
		fmt.Printf("%s escaped!\n", params[0])
	}

	return nil
}
