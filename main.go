package main

import (
	"apii_v1/handlers"
	"apii_v1/init_"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func init() {
	init_.ConnectDb()
	init_.AutoMigrate()
}

func main() {
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Static("/static", "./static")

	app.Get("/", handlers.Home)

	log.Fatal(app.Listen(":3000"))
}

// user := models.Worker{Name: "Janis", Position: "Manager", IsAdmin: true}
// init_.DB.Create(&user)
// fmt.Println("User created")
