package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	book "github.com/sathish-30/Golang-fiber/books"
	"github.com/sathish-30/Golang-fiber/database"
	"github.com/sathish-30/Golang-fiber/routes"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("Database Connection successfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("DataBase Migrated")
}

func main() {

	// Creates a new fiber app instance
	app := fiber.New()
	initDatabase()
	// Setting up routes
	routes.SetUpRoutes(app)

	app.Listen(":3000")

}
