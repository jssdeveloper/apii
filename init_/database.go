package init_

import (
	"apii_v1/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("apii_db.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to sqlite db")
		panic(err)
	}
	fmt.Println("Connected to db")
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Worker{})
	fmt.Println("Automigrated")
}
