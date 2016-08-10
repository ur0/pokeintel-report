package models

type Encounter struct {
	PokemonId     uint32  `json:"pokemon_id"`
	EncounterTime string  `json:"encounter_time"` // A RFC3339 formatted timestamp
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}
