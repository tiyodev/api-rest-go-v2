package models

// Specy entity
type Specy struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `gorm:"column:average_height" json:"average_height"`
	SkinColors      string   `gorm:"column:skin_colors" json:"skin_colors"`
	HairColors      string   `gorm:"column:hair_colors" json:"hair_colors"`
	EyeColors       string   `gorm:"column:eye_colors" json:"eye_colors"`
	AverageLifespan string   `gorm:"column:average_lifespan" json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	People          []People `gorm:"many2many:people_species;association_jointable_foreignkey:people;jointable_foreignkey:species" json:"peoples"`
	Films           []Film   `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	URL             int      `json:"url"`
	ID              uint     `gorm:"primary_key" json:"id"`
}
