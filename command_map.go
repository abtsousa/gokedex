package main

import (
	"fmt"
)

func commandMapN(cfg *Config, params ...string) error {
	return commandMap(false, cfg)
}

func commandMapP(cfg *Config, params ...string) error {
	if cfg.prev == nil {
		fmt.Println("You're already on the first page.")
		return nil
	}
	return commandMap(true, cfg)
}

func commandMap(goBack bool, cfg *Config) error {

	var url *string
	if goBack {
		url = cfg.prev
	} else {
		url = cfg.next
	}

	loc, err := cfg.client.ListLocations(url, cfg.cache)
	if err != nil {
		return err
	}

	cfg.next = loc.Next
	cfg.prev = loc.Previous

	for _, r := range loc.Results {
		fmt.Println(r.Name)
	}
	// fmt.Printf("***DEBUG:\nloc.Next: %v\nloc.Previous: %v\ncfg.next: %v\n cfg.prev: %v", loc.Next, loc.Previous, cfg.next, cfg.prev)
	return nil
}
