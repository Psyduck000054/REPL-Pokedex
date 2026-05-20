package pokeapi

type PokemonName struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokeBall struct {
	Container PokemonName `json:"pokemon"`
}

type AreaEncounterInfo struct {
	Results []PokeBall `json:"pokemon_encounters"`
}
