package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sathish-30/Golang-fiber/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book Book
	book.Title = "Think like a Monk"
	book.Author = "Jay Shetty"
	book.Rating = 9
	db.Create(&book)

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book found with the given ID")
	} else {
		db.Delete(&book)
		return c.SendString("Book succesfully deleted")
	}

}
