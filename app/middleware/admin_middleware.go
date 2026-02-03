package middleware

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AdminProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {

		isAdmin := utils.IsAdmin(c)

		if !isAdmin {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": true,
				"msg":   "Forbidden: Admin privileges required",
			})
		}

		return c.Next()
	}
}
