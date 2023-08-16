package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single Book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("Add new Book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
