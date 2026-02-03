package routes

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func bookRouter(router fiber.Router) {
	router.Get("/", handlers.GetBooks)
	router.Post("/", handlers.CreateBook)
	router.Get("/authors", handlers.GetBooksWithAuthor)
}
