package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sathish-30/Golang-fiber/database"
	"github.com/sathish-30/Golang-fiber/routes"
	_ "gorm.io/driver/sqlite"
)

func main() {

	// Creates a new fiber app instance
	app := fiber.New()
	//Connect with the database
	database.Connect()
	// Setting up routes
	router.SetUpRoutes(app)

	app.Listen(":3000")

}
