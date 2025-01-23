package main

import (
	"fmt"
)

func commandHelp(cfg *Config, params ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for i := range mp {
		fmt.Printf("%v: %v\n", mp[i].name, mp[i].description)
	}
	return nil
}
