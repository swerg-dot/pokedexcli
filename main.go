package main

import (
	"fmt"
	"time"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon 	    map[string]pokeapi.Pokemon
}

func main() {
	// Initialize the cache with a 1-hour expiration interval
	cache := pokecache.NewCache(time.Hour)

	// Create a new pokeapi client, passing the cache into the client constructor
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour), // The pokeapi client already has access to the cache
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	// Start the REPL
	startRepl(&cfg)

	// Example for cache usage
	fmt.Println("Cache initialized:", cache)
}
