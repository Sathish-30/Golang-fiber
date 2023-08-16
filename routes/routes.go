package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sathish-30/Golang-fiber/books"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/books", book.NewBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)
}
