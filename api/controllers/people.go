package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tiyodev/api-rest-go-v1/api/models"
	"github.com/tiyodev/api-rest-go-v1/api/responses"
)

// GetPeople : get people by id
func (server *Server) GetPeople(resWriter http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(resWriter, http.StatusBadRequest, err)
		return
	}

	people, err := models.FindPeopleByID(server.DB, pid)

	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(resWriter, http.StatusOK, people)
}

// GetPeoples : get all peoples with optional limit and offset
func (server *Server) GetPeoples(resWriter http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	limit, err := strconv.ParseUint(params.Get("limit"), 10, 64)
	offset, err := strconv.ParseUint(params.Get("offset"), 10, 64)

	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}

	peoples, err := models.FindAllUsers(server.DB, limit, offset)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(resWriter, http.StatusOK, peoples)
}

// CreatePeople : create a new people
func (server *Server) CreatePeople(resWriter http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	fmt.Println(body)
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	people := models.People{}

	// Generate default value
	people.Prepare()

	// init body with json body
	err = json.Unmarshal(body, &people)
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	// Get last people ID
	lastID, err := models.FindLastPeopleID(server.DB)

	// Increment ID and URL
	people.ID = lastID + 1
	people.URL = lastID + 1

	// Validate required values
	err = people.Validate()
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	// Save new people in the DB
	peopleCreated, err := people.SavePeople(server.DB)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}

	// Get created people
	peopleCreated, err = models.FindPeopleByID(server.DB, peopleCreated.ID)

	// Return created people
	responses.JSON(resWriter, http.StatusCreated, peopleCreated)
}

// UpdatePeople : Update people informations
func (server *Server) UpdatePeople(resWriter http.ResponseWriter, req *http.Request) {

	// Get URL ID parameter
	vars := mux.Vars(req)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(resWriter, http.StatusBadRequest, err)
		return
	}

	// Parse all body parameter
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	// Find the person to be modified
	var people *models.People
	people, err = models.FindPeopleByID(server.DB, uid)

	// Set new values
	err = json.Unmarshal(body, &people)
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	// Validate people
	err = people.Validate()
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	// Update new people values
	updatedUser, err := people.UpdatePeople(server.DB, uid)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(resWriter, http.StatusOK, updatedUser)
}

// DeletePeople : remove people
func (server *Server) DeletePeople(resWriter http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	people := models.People{}

	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(resWriter, http.StatusBadRequest, err)
		return
	}

	_, err = people.DeletePeople(server.DB, uid)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(resWriter, http.StatusOK, "ok")
}
