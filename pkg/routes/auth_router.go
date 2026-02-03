package routes

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func authRouter(router fiber.Router) {
	router.Post("/auth/login", handlers.Login)
}
