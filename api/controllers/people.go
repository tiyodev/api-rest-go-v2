package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tiyodev/api-rest-go-v2/api/models"
	"github.com/tiyodev/api-rest-go-v2/api/responses"
)

// GetPeople : get people by id
func (server *Server) GetPeople(resWriter http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	// Get URL ID parameter
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
	var limit uint64
	var offset uint64
	var err error

	// Get Limit and Offset optional URL params
	params := req.URL.Query()

	if l := params.Get("limit"); l != "" {
		limit, err = strconv.ParseUint(params.Get("limit"), 10, 64)
	}

	if o := params.Get("offset"); o != "" {
		offset, err = strconv.ParseUint(params.Get("offset"), 10, 64)
	}

	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}

	peoples, err := models.FindAllPeoples(server.DB, limit, offset)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(resWriter, http.StatusOK, peoples)
}

// CreatePeople : create a new people
func (server *Server) CreatePeople(resWriter http.ResponseWriter, req *http.Request) {
	// Get all HTTP body params
	body, err := ioutil.ReadAll(req.Body)
	fmt.Println(body)
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	people := models.People{}

	// Generate default value
	people.Prepare()

	// Init people with HTTP body params
	err = json.Unmarshal(body, &people)
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	// Validate required values
	err = people.Validate()
	if err != nil {
		responses.ERROR(resWriter, http.StatusUnprocessableEntity, err)
		return
	}

	// Save new people in the DB
	err = people.SavePeople(server.DB)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}

	// URL must be the same as ID
	people.URL = people.ID

	// Update created people with URL as same as ID
	err = people.UpdatePeople(server.DB, people.ID)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}

	// Return created people
	responses.JSON(resWriter, http.StatusCreated, people)
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
	people, err := models.FindPeopleByID(server.DB, uid)

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
	err = people.UpdatePeople(server.DB, uid)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(resWriter, http.StatusOK, people)
}

// DeletePeople : remove people
func (server *Server) DeletePeople(resWriter http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// Get URL ID parameter
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(resWriter, http.StatusBadRequest, err)
		return
	}

	_, err = models.DeletePeople(server.DB, uid)
	if err != nil {
		responses.ERROR(resWriter, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(resWriter, http.StatusOK, "ok")
}
