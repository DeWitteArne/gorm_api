package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

//CreatePerson gets a body from a call to create a new Person
func CreatePerson(w http.ResponseWriter, r *http.Request) error {

	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}
	defer r.Body.Close()
	fmt.Println(person.CarID)
	answer, err := models.NewPerson(person.Firstname, person.Lastname, person.CarID, person.Username)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, answer)
	}

	return writeJSON(w, http.StatusOK, answer)
}

//GetPersons asks all the persons and returns them
func GetPersons(w http.ResponseWriter, r *http.Request) error {

	persons, err := models.ReadAllPersons()

	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}

	return writeJSON(w, http.StatusOK, persons)

}

//GetPerson asks a person with a specific id and returns the value
func GetPerson(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	person, err := models.ReadOnePerson(id)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}
	return writeJSON(w, http.StatusOK, person)
}

//ModifyPerson sends a body and id of a person to update and returns an answer
func ModifyPerson(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}
	defer r.Body.Close()

	answer, err := models.UpdatePerson(id, person)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}

	return writeJSON(w, http.StatusOK, answer)
}

//RemovePerson asks to delete the person with the specific id and returns an answer
func RemovePerson(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	answer, err := models.DeletePerson(id)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}
	return writeJSON(w, http.StatusOK, answer)
}
