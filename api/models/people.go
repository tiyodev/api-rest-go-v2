package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// People entity
type People struct {
	Name        string     `json:"name"`
	Height      string     `json:"height"`
	Mass        string     `json:"mass"`
	HairColor   string     `json:"hair_color"`
	SkinColor   string     `json:"skin_color"`
	EyeColor    string     `json:"eye_color"`
	BirthYear   string     `json:"birth_year"`
	Gender      string     `json:"gender"`
	Homeworld   Planet     `gorm:"foreignkey:HomeworldID;" json:"homeworld"`
	HomeworldID uint       `gorm:"column:homeworld" json:"homeworld_id"`
	Films       []Film     `gorm:"many2many:films_people;association_jointable_foreignkey:films;jointable_foreignkey:people" json:"films"`
	Species     []Specy    `gorm:"many2many:people_species;association_jointable_foreignkey:species;jointable_foreignkey:people" json:"species"`
	Vehicles    []Vehicle  `gorm:"many2many:people_vehicles;association_jointable_foreignkey:vehicles;jointable_foreignkey:people" json:"vehicles"`
	Starships   []Starship `gorm:"many2many:people_starships;association_jointable_foreignkey:starships;jointable_foreignkey:people" json:"starships"`
	Created     string     `json:"created"`
	Edited      string     `json:"edited"`
	URL         uint64     `json:"url"`
	ID          uint64     `gorm:"primary_key;auto_increment" json:"id"`
}

// TableName Get table name
func (People) TableName() string {
	return "people"
}

// Prepare set default value
func (people *People) Prepare() {
	people.Height = "unknown"
	people.Name = "unknown"
	people.Mass = "unknown"
	people.HairColor = "unknown"
	people.SkinColor = "unknown"
	people.EyeColor = "unknown"
	people.BirthYear = "unknown"
	people.Gender = "na"
	people.Homeworld = Planet{}
	people.Created = time.Now().String()
	people.Edited = time.Now().String()
	people.URL = 0
	people.ID = 0
}

// Validate require people values
func (people *People) Validate() error {
	if people.Name == "" {
		return errors.New("Required Name")
	}
	if people.HomeworldID < 1 {
		return errors.New("Required Homeworld")
	}
	if people.URL < 1 {
		return errors.New("Required URL")
	}
	if people.ID < 1 {
		return errors.New("Required ID")
	}

	if people.ID != people.URL {
		return errors.New("ID and URL must be equal")
	}
	return nil
}

// FindLastPeopleID find People by id
func FindLastPeopleID(db *gorm.DB) (uint64, error) {
	var people People
	err := db.Debug().Select("cast(id as int) id").Order("id DESC").Model(People{}).First(&people).Error
	return people.ID, err
}

// FindPeopleByID find People by id
func FindPeopleByID(db *gorm.DB, uid uint64) (*People, error) {
	// Preload people relations and get first people by ID
	var people People
	err := db.Debug().Preload("Homeworld").Preload("Films").Preload("Vehicles").Preload("Starships").Preload("Species").Model(People{}).Where("id = ?", uid).Take(&people).Error
	if err != nil {
		return &People{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &People{}, errors.New("People Not Found")
	}
	return &people, err
}

// FindAllPeoples find all peoples
func FindAllPeoples(db *gorm.DB, limit uint64, offset uint64) (*[]People, error) {
	var peoples []People

	// Preload people relations and get all people
	query := db.Debug().Preload("Homeworld").Preload("Films").Preload("Vehicles").Preload("Starships").Preload("Species").Model(People{})

	// Add limit filter if needed
	if limit > 0 {
		query = query.Limit(limit)
	}

	// Add offset filter if needed
	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&peoples).Error
	if err != nil {
		return &[]People{}, err
	}
	return &peoples, err
}

// SavePeople create people
func (people *People) SavePeople(db *gorm.DB) (*People, error) {
	var err error

	err = db.Debug().Create(&people).Error
	if err != nil {
		return &People{}, err
	}
	return people, nil
}

// UpdatePeople : Update people
func (people *People) UpdatePeople(db *gorm.DB, uid uint64) (*People, error) {

	fmt.Println(people)

	db = db.Debug().Model(&people).UpdateColumns(
		map[string]interface{}{
			"HomeworldID": people.HomeworldID,
			"name":        people.Name,
			"height":      people.Height,
			"mass":        people.Mass,
			"hair_color":  people.HairColor,
			"skin_color":  people.SkinColor,
			"eye_color":   people.EyeColor,
			"birth_year":  people.BirthYear,
			"gender":      people.Gender,
			"edited":      time.Now(),
		},
	)
	if db.Error != nil {
		return &People{}, db.Error
	}
	// This is the display of updated user
	errors := db.Debug().Model(&People{}).Where("id = ?", uid).Take(&people).Error
	if errors != nil {
		return &People{}, errors
	}
	return people, nil
}

// DeletePeople : Delete a people
func (people *People) DeletePeople(db *gorm.DB, uid uint64) (int64, error) {

	db = db.Debug().Model(&People{}).Where("id = ?", uid).Take(&people).Delete(&People{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
