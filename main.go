package main

import (
	"time"

	"github.com/abtsousa/gokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		client: client,
	}
	startRepl(cfg)
}
