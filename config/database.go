package config

import (
	"apii_v1/entities"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to sqlite db")
		panic(err)
	}
	fmt.Println("Connected to db")
	db.AutoMigrate(&entities.Worker{})
	fmt.Println("Automigrated")

	return nil
}
