package models

type Encounter struct {
  PokemonId uint32 `json:"pokemon_id"`
  EncounterTime uint64 `json:"encounter_time"`
  Latitude float64
  Longitude float64  
}
