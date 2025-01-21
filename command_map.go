package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapN(cfg *Config) error {
	return commandMap(false, cfg)
}

func commandMapP(cfg *Config) error {
	return commandMap(true, cfg)
}

func commandMap(goBack bool, cfg *Config) error {
	type Locations struct {
		Count    int     `json:"count"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	url := "https://pokeapi.co/api/v2/location-area/"

	if goBack {
		if cfg.prev != "" {
			url = cfg.prev
		} else {
			fmt.Println("You're on the first page.")
			return nil
		}
	} else {
		if cfg.next != "" {
			url = cfg.next
		}
	}
	rsp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("couldn't get response: %v", err)
	}

	dat, err := io.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("couldn't parse received data: %v", err)
	}

	loc := Locations{}
	err = json.Unmarshal(dat, &loc)
	if err != nil {
		return fmt.Errorf("couldn't unmarshal received data: %v", err)
	}

	if loc.Next != nil {
		cfg.next = *loc.Next
	} else {
		cfg.next = ""
	}
	if loc.Previous != nil {
		cfg.prev = *loc.Previous
	} else {
		cfg.prev = ""
	}
	for _, r := range loc.Results {
		fmt.Println(r.Name)
	}

	fmt.Printf("***DEBUG*** cfg.prev: %s\tcfg.next: %s\n", cfg.prev, cfg.next)
	return nil
}
