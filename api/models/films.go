package models

// Film entity
type Film struct {
	Title        string     `json:"title"`
	EpisodeID    int        `json:"episode_id"`
	OpeningCrawl string     `json:"opening_crawl"`
	Director     string     `json:"director"`
	Producer     string     `json:"producer"`
	Characters   []People   `gorm:"many2many:films_people;association_jointable_foreignkey:people;jointable_foreignkey:films" json:"characters"`
	Planets      []Planet   `json:"planets"`
	Starships    []Starship `json:"starships"`
	Vehicles     []Vehicle  `json:"vehicles"`
	Species      []Specy    `json:"species"`
	ReleaseDate  string     `json:"release_date"`
	Dreated      string     `json:"dreated"`
	Edited       string     `json:"edited"`
	URL          int        `json:"url"`
	ID           uint       `gorm:"primary_key" json:"id"`
}
