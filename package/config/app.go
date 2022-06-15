package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func connectDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=restaurant password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
}
