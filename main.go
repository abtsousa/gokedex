package main

import (
	"time"

	"github.com/abtsousa/gokedex/internal/cache"
	"github.com/abtsousa/gokedex/internal/pokeapi"
)

type pokedex = map[string]pokeapi.Pokemon

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cache := cache.NewCache(5 * time.Second)
	pokedex := pokedex{}
	cfg := &Config{
		client:  client,
		cache:   &cache,
		pokedex: pokedex,
	}
	startRepl(cfg)
}
