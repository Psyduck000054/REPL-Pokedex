package main

import (
	"pokedex/internal/pokeapi"
)

type config struct {
	pokeClient *pokeapi.Client
	pokedex    map[string]pokeapi.Pokemon
	previous   *string
	next       *string
}
