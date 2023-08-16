package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sathish-30/Golang-fiber/handler"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/api/v1/books", handler.GetBooks)
	app.Get("/api/v1/book/:id", handler.GetBook)
	app.Post("/api/v1/handler", handler.NewBook)
	app.Delete("/api/v1/book/:id", handler.DeleteBook)
}
