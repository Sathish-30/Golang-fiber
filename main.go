package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sathish-30/Golang-fiber/database"
	"github.com/sathish-30/Golang-fiber/routes"
	_ "gorm.io/driver/sqlite"
	"log"
	"os"
)

func main() {

	// Creates a new fiber app instance
	app := fiber.New()
	//Connect with the database
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	database.Connect(config)
	// Setting up routes
	router.SetUpRoutes(app)

	app.Listen(":3000")

}
