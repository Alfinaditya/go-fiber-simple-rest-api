package routes

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func userRouter(router fiber.Router) {
	router.Post("/logout", handlers.Logout)
	router.Get("/me", handlers.Me)
}
