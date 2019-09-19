package models_test

import (
	"testing"

	"github.com/tiyodev/api-rest-go-v1/api/models"
)

func TestTableName(t *testing.T) {

	var people models.People

	result := people.TableName()

	if result != "people" {
		t.Errorf("TableName(\"\") failed, expected %v, got %v", "people", result)
	}
}

func TestPrepare(t *testing.T) {
	people := models.People{}
	people.Prepare()

	if people.Height != "unknown" {
		t.Errorf("Prepare(\"\") failed, expected %v, got %v", "unknown", people.Height)
	}

	if people.Gender != "na" {
		t.Errorf("Prepare(\"\") failed, expected %v, got %v", "na", people.Gender)
	}

	if people.URL != 0 {
		t.Errorf("Prepare(\"\") failed, expected %v, got %v", "> 0", people.URL)
	}

	if people.ID != 0 {
		t.Errorf("Prepare(\"\") failed, expected %v, got %v", "> 0", people.ID)
	}
}

func TestValidateFailed(t *testing.T) {
	people := models.People{
		Gender:      "male",
		HairColor:   "grey",
		HomeworldID: 0,
	}
	err := people.Validate()

	if err == nil {
		t.Errorf("Validate(\"\") failed, expected %v, got %v", "Required Name", err)
	}

	people = models.People{
		Gender:      "male",
		Name:        "Tony",
		HomeworldID: 0,
	}
	err = people.Validate()

	if err == nil {
		t.Errorf("Validate(\"\") failed, expected %v, got %v", "Required Homeworld", err)
	}

	people = models.People{
		Gender:      "male",
		Name:        "Tony",
		HomeworldID: 1,
	}
	err = people.Validate()

	if err == nil {
		t.Errorf("Validate(\"\") failed, expected %v, got %v", "Required URL", err)
	}

	people = models.People{
		Gender:      "male",
		Name:        "Tony",
		HomeworldID: 1,
		URL:         2,
	}
	err = people.Validate()

	if err == nil {
		t.Errorf("Validate(\"\") failed, expected %v, got %v", "Required ID", err)
	}

	people = models.People{
		Gender:      "male",
		Name:        "Tony",
		HomeworldID: 1,
		URL:         2,
		ID:          3,
	}
	err = people.Validate()

	if err == nil {
		t.Errorf("Validate(\"\") failed, expected %v, got %v", "ID and URL must be equal", err)
	}

	people = models.People{
		Gender:      "male",
		Name:        "Tony",
		HomeworldID: 1,
		URL:         2,
		ID:          2,
	}
	err = people.Validate()

	if err != nil {
		t.Errorf("Validate(\"\") failed, expected %v, got %v", nil, err)
	}
}

func TestValidateSuccess(t *testing.T) {
	people := models.People{
		Gender:      "male",
		Name:        "Tony",
		HomeworldID: 1,
		URL:         2,
		ID:          2,
	}
	err := people.Validate()

	if err != nil {
		t.Errorf("Validate(\"\") failed, expected %v, got %v", nil, err)
	}
}
