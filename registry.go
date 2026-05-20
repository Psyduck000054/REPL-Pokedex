package main

// ---------------------------------------------------------
// COMMAND MAPPING
// ---------------------------------------------------------

type cliCommand struct {
	name     string
	desc     string
	callback func(c *config, param1 string) error
}

var commandList = map[string]cliCommand{
	"exit": {
		name:     "exit",
		desc:     "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name:     "help",
		desc:     "Return a tutorial",
		callback: commandHelp,
	},
	"map": {
		name:     "map",
		desc:     "List the locations of the next page",
		callback: commandMap,
	},
	"mapb": {
		name:     "mapb",
		desc:     "List the locations of the previous page",
		callback: commandMapb,
	},
	"explore": {
		name:     "explore",
		desc:     "List the Pokemons that can be encountered in an area",
		callback: commandExplore,
	},
	"catch": {
		name:     "catch",
		desc:     "Attempt to catch a Pokemon!",
		callback: commandCatch,
	},
	"check": {
		name:     "check",
		desc:     "Check the details of a Pokemon in your Pokedex",
		callback: commandCheck,
	},
	"inspect": {
		name:     "inspect",
		desc:     "Check the details of a Pokemon in your Pokedex",
		callback: commandInspect,
	},
	"pokedex": {
		name:     "pokedex",
		desc:     "List all Pokemon in your Pokedex",
		callback: commandPokedex,
	},
}
