package main

import (
	"time"

	"github.com/abtsousa/gokedex/internal/cache"
	"github.com/abtsousa/gokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cache := cache.NewCache(5 * time.Second)
	cfg := &Config{
		client: client,
		cache:  &cache,
	}
	startRepl(cfg)
}
