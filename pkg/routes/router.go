package routes

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router) {
	authRouter(router)

	protectedRoot := router.Group("/root", middleware.AuthProtected(), middleware.AdminProtected())
	rootRouter(protectedRoot)

	protectedBook := router.Group("/books", middleware.AuthProtected())
	protectedAuthor := router.Group("/authors", middleware.AuthProtected())
	protectedUser := router.Group("/users", middleware.AuthProtected())

	userRouter(protectedUser)
	bookRouter(protectedBook)
	authorRouter(protectedAuthor)
}
