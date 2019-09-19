package models

// Vehicle entity
type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 int      `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	Pilots               []People `gorm:"many2many:people_vehicles;association_jointable_foreignkey:people;jointable_foreignkey:vehicles" json:"pilots"`
	Films                []Film   `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  int      `json:"url"`
	ID                   uint     `gorm:"primary_key" json:"id"`
}
