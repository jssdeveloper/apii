package main

import (
	"apii_v1/handlers"
	"apii_v1/init_"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	init_.ConnectDb()
	init_.AutoMigrate()
}

func main() {
	app := fiber.New()

	app.Get("/", handlers.Home)

	log.Fatal(app.Listen(":3000"))
}

// user := models.Worker{Name: "Janis", Position: "Manager", IsAdmin: true}
// init_.DB.Create(&user)
// fmt.Println("User created")
