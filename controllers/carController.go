package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

type httpAPIFunc func(w http.ResponseWriter, r *http.Request) error

//MakeHTTPHandler handles the api calls
func MakeHTTPHandler(f httpAPIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

//GetCars calls de car struct en gets all the cars from the database and returns them into json
func GetCars(w http.ResponseWriter, r *http.Request) error {
	cars, err := models.ReadAllCars()
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}
	return writeJSON(w, http.StatusOK, cars)

}

//CreateCar gets a body through a post request and sends a Car struct to the function CreateNewCar from car.go
func CreateCar(w http.ResponseWriter, r *http.Request) error {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		return err
	}
	defer r.Body.Close()
	answer, err := models.NewCar(car.Licenseplate, car.Color, car.Doors, car.Length)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}

	return writeJSON(w, http.StatusOK, answer)

}

//GetCar function that gets id from url, asks car by id and returns car
func GetCar(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	car, err := models.ReadOneCar(id)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	} else {
		return writeJSON(w, http.StatusOK, car)
	}
}

//ModifyCar receives a body and id, both passed by to car.go to update the car with this specific id
func ModifyCar(w http.ResponseWriter, r *http.Request) error {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		return err
	}
	defer r.Body.Close()
	params := mux.Vars(r)
	id := params["id"]
	answer, err := models.UpdateCar(id, car)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	} else {
		return writeJSON(w, http.StatusOK, answer)
	}

}

//RemoveCar gets an id and passes it to the DeleteCar function wich gives an answer and/or err back
func RemoveCar(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	answer, err := models.DeleteCar(id)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	} else {
		return writeJSON(w, http.StatusOK, answer)
	}
}

func writeJSON(w http.ResponseWriter, i int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(i)
	return json.NewEncoder(w).Encode(v)
}
