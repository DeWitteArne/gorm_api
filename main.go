package main

import (
	"log"
	"net/http"

	"./controllers"
	"./data"
	"./models"
	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := data.GetDb()
	db.AutoMigrate(&models.Car{})
	db.AutoMigrate(&models.Person{})
	r := mux.NewRouter()
	r.HandleFunc("/cars", controllers.MakeHTTPHandler(controllers.GetCars)).Methods("GET")
	r.HandleFunc("/cars", controllers.MakeHTTPHandler(controllers.CreateCar)).Methods("POST")
	r.HandleFunc("/cars/{id}", controllers.MakeHTTPHandler(controllers.GetCar)).Methods("GET")
	r.HandleFunc("/cars/{id}", controllers.MakeHTTPHandler(controllers.ModifyCar)).Methods("PUT")
	r.HandleFunc("/cars/{id}", controllers.MakeHTTPHandler(controllers.RemoveCar)).Methods("DELETE")
	r.HandleFunc("/persons", controllers.MakeHTTPHandler(controllers.GetPersons)).Methods("GET")
	r.HandleFunc("/persons", controllers.MakeHTTPHandler(controllers.CreatePerson)).Methods("POST")
	r.HandleFunc("/persons/{id}", controllers.MakeHTTPHandler(controllers.GetPerson)).Methods("GET")
	r.HandleFunc("/persons/{id}", controllers.MakeHTTPHandler(controllers.ModifyPerson)).Methods("PUT")
	r.HandleFunc("/persons/{id}", controllers.MakeHTTPHandler(controllers.RemovePerson)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
