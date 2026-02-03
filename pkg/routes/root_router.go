package routes

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func rootRouter(router fiber.Router) {
	router.Post("/revoke-token/:id", handlers.RevokeTokenVersion)
	router.Get("/users", handlers.GetUsers)
}
