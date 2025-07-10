package main

import (
	"time"

	"github.com/wildwinds86/pokedex/internal"
)

var pokecache internal.Cache

func main() {

	pokeClient := NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	pokecache = internal.NewCache(5 * time.Second)
	go pokecache.ReapLoop()

	startRepl(cfg)
}
