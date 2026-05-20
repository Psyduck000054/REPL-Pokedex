package pokeapi

type StatInfo struct {
	Name string `json:"name"`
}

type Stat struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

type Type struct {
	TypeInfo TypeInfo `json:"type"`
}

type Pokemon struct {
	Name   string `json:"name"`
	BaseXP int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []Stat `json:"stats"`
	Types  []Type `json:"types"`
}
