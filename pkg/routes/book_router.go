package routes

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func bookRouter(router fiber.Router) {
	router.Get("/", handlers.GetBooks)
	router.Get("/authors", handlers.GetBooksWithAuthor)
	// router.Get("/books/:id", getBook)
	// router.Post("/books", createBook)
	// router.Put("/books/:id", updateBook)
	// router.Delete("/books/:id", deleteBook)
}
