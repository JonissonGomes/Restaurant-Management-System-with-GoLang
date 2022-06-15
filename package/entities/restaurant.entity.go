package entities

import (
	"github.com/JonissonGomes/Restaurant-Management-System-with-GoLang/package/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Restaurant struct {
	gorm.Model
	Name          string `json:"name"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	Instragram    string `json:"instagram"`
	QuantityStars int    `json:"quantityStars"`
}

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	db.AutoMigrate(&Restaurant{})
}

func (r *Restaurant) CreateRestaurant() *Restaurant {
	db.NewRecord(r)
	db.Create(&r)
	return r
}

func GetAllRestaurant() []Restaurant {
	var restaurants []Restaurant
	db.Find(&restaurants)
	return restaurants
}

func GetRestaurantById(id int64) (*Restaurant, *gorm.DB) {
	var restaurant Restaurant
	db := db.Where("id = ?", id).Find(&restaurant)
	return &restaurant, db
}

func DeleteRestaurant(id int64) *gorm.DB {
	db := db.Where("id = ?", id).Delete(&Restaurant{})
	return db
}
