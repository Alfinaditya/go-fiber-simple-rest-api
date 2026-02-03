package server

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func StartServer() {

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")

	routes.AddRoutes(api)

	log.Fatal(app.Listen(":3001"))
}
