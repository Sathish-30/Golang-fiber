package handler

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"github.com/sathish-30/Golang-fiber/database"
	"github.com/sathish-30/Golang-fiber/model"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book found with the given ID")
	} else {
		db.Delete(&book)
		return c.SendString("Book succesfully deleted")
	}

}
