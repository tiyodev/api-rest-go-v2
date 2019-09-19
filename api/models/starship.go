package models

// Starship entity
type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"mglt"`
	StarshipClass        string   `json:"starship_class"`
	Pilots               []People `gorm:"many2many:people_starships;association_jointable_foreignkey:people;jointable_foreignkey:starships" json:"pilots"`
	Films                []Film   `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  int      `json:"url"`
	ID                   int      `gorm:"primary_key" json:"id"`
}
