package main

import (
	"fmt"
)

func commandInspect(cfg *Config, params ...string) error {

	if len(params) == 0 {
		return fmt.Errorf("You must provide a valid Pokémon name.")
	}

	pkm, bool := cfg.pokedex[params[0]]

	if !bool {
		fmt.Printf("You have not yet caught %s\n", params[0])
		return nil
	}

	// Name
	fmt.Printf("Name: %s\n", pkm.Name)
	// Height
	fmt.Printf("Height: %d\n", pkm.Height)
	// Stats
	fmt.Println("Stats:")
	for _, stat := range pkm.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	// Types
	fmt.Println("Types:")
	for _, typ := range pkm.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}

	return nil
}
