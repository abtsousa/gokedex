package main

import "fmt"

func commandExplore(cfg *Config, params ...string) error {

	if len(params) == 0 {
		return fmt.Errorf("You must provide a valid location name.")
	}

	det, err := cfg.client.ListLocationDetails(params[0], cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", params[0])

	if len(det.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found!")
		return nil
	}

	fmt.Println("Found Pokémon:")
	for _, pkm := range det.PokemonEncounters {
		fmt.Printf("- %s\n", pkm.Pokemon.Name)
	}
	return nil
}
