package data

import (
	"github.com/jinzhu/gorm"
)

// GetDb function that returns the database
func GetDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@/products?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
