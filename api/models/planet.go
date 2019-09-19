package models

// Planet entity
type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `gorm:"column:rotation_period" json:"rotation_period"`
	OrbitalPeriod  string   `gorm:"column:orbital_period" json:"orbital_period"`
	Diameter       string   `json:"Diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `gorm:"column:surface_water" json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []People `gorm:"foreignkey:Homeworld;association_foreignkey:ID"`
	Films          []Film   `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            int      `json:"url"`
	ID             uint     `gorm:"primary_key"`
}
