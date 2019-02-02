package models

import (
	"../data"
	"github.com/jinzhu/gorm"
)

//Person representing a person
type Person struct {
	gorm.Model
	Firstname string
	Lastname  string
	Car       Car `gorm:"foreignkey:CarID"`
	CarID     uint
	Username  string
}

//NewPerson creates a new person and creates a record for that person in the database
func NewPerson(firstname string, lastname string, carID uint, username string) (string, error) {
	db := data.GetDb()

	var car Car
	if err := db.First(&car, carID).Error; err != nil {
		return "Car failed", err
	}
	person := Person{Firstname: firstname, Lastname: lastname, Car: car, Username: username}
	if err := db.Create(&person).Error; err != nil {
		return "Failed", err
	}

	return "Success", nil
}

//ReadAllPersons returns all persons in the database
func ReadAllPersons() ([]Person, error) {
	db := data.GetDb()
	var persons []Person

	if err := db.Find(&persons).Error; err != nil {
		return nil, err
	}
	for i, person := range persons {
		var c Car
		if err := db.First(&c, person.CarID).Error; err != nil {
			return nil, err
		}
		persons[i].Car = c
	}

	return persons, nil
}

//ReadOnePerson returns one person with the givin id
func ReadOnePerson(id string) (Person, error) {
	db := data.GetDb()
	var person Person
	if err := db.First(&person).Error; err != nil {
		return Person{}, err
	}
	var car Car
	if err := db.First(&car, person.CarID).Error; err != nil {
		return Person{}, err
	}
	person.Car = car
	return person, nil

}

//UpdatePerson updates a person with the specific id in the database
func UpdatePerson(id string, person Person) (string, error) {
	db := data.GetDb()
	var p Person
	if err := db.First(&p, id).Error; err != nil {
		return "Failed", err
	}

	if person.Firstname != "" {
		p.Firstname = person.Firstname
	}
	if person.Lastname != "" {
		p.Lastname = person.Lastname
	}

	if person.Username != "" {
		p.Username = person.Username
	}

	if person.CarID != 0 {
		var car Car
		if err := db.First(&car, person.CarID).Error; err != nil {
			return "Failed", err
		}
		p.CarID = person.CarID
		p.Car = car
	}

	if err := db.Save(&p).Error; err != nil {
		return "Failed", err
	}
	return "Success", nil
}

//DeletePerson deletes a person from the database with the specific id
func DeletePerson(id string) (string, error) {
	db := data.GetDb()
	var person Person
	if err := db.First(&person, id).Error; err != nil {
		return "Failed", err
	}
	if err := db.Delete(&person).Error; err != nil {
		return "Failed", err
	}
	return "Success", nil
}
