package models

import (
	"time"

	"../data"
	"github.com/jinzhu/gorm"
)

// Car represents a Car
type Car struct {
	gorm.Model
	Licenseplate string
	Color        string
	Manufatured  time.Time
	Doors        int
	Length       float64
}

// NewCar function that creates a new car and stores it in the database
func NewCar(licenseplate string, color string, doors int, length float64) (string, error) {
	db := data.GetDb()
	car := Car{Licenseplate: licenseplate, Color: color, Doors: doors, Length: length, Manufatured: time.Now()}
	if err := db.Create(&car).Error; err != nil {
		return "Failed", err
	} else {
		return "Success", nil
	}

}

//CreateNewCar creates a car and gets as parameter an Car struct
func CreateNewCar(car Car) string {
	db := data.GetDb()
	db.Create(car)

	return "done"
}

// ReadAllCars is a function that returns all cars in the database
func ReadAllCars() ([]Car, error) {
	db := data.GetDb()
	var cars []Car
	if err := db.Find(&cars).Error; err != nil {
		return nil, err
	} else {
		return cars, nil
	}

}

// ReadOneCar returns a car with the given id
func ReadOneCar(id string) (Car, error) {
	db := data.GetDb()
	var car Car
	if err := db.First(&car, id).Error; err != nil {
		return Car{}, err
	} else {
		return car, nil
	}
}

// UpdateCar modifies a car with the given id
func UpdateCar(id string, car Car) (string, error) {
	db := data.GetDb()
	var carUpdate Car
	db.First(&carUpdate, id)
	if car.Color != "" {
		carUpdate.Color = car.Color
	}
	if car.Doors != 0 {
		carUpdate.Doors = car.Doors
	}
	if car.Length != 0 {
		carUpdate.Length = car.Length
	}
	if car.Licenseplate != "" {
		carUpdate.Licenseplate = car.Licenseplate
	}
	if err := db.Save(&carUpdate).Error; err != nil {
		return "Failed", err
	} else {
		return "Success", nil
	}

}

//DeleteCar function that deletes a car and returns done when finished
func DeleteCar(id string) (string, error) {
	db := data.GetDb()
	var carDelete Car
	db.First(&carDelete, id)
	if err := db.Delete(&carDelete).Error; err != nil {
		return "Failed", err
	} else {
		return "Success", nil
	}

}
