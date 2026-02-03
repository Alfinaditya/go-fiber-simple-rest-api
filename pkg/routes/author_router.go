package routes

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func authorRouter(router fiber.Router) {
	router.Get("/", handlers.GetAuthors)
	router.Post("/", handlers.CreateAuthor)
	router.Put("/:id", handlers.UpdateAuthor)
	router.Delete("/:id", handlers.DeleteAuthor)
}
