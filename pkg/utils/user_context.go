package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTokenVersion(c *fiber.Ctx) int {
	val, ok := c.Locals("TokenVersion").(int)
	if !ok {
		return 0
	}
	return val
}

func GetUserID(c *fiber.Ctx) string {
	val, ok := c.Locals("UserID").(string)
	if !ok {
		return ""
	}

	if _, err := uuid.Parse(val); err != nil {
		return ""
	}
	return val
}

func IsAdmin(c *fiber.Ctx) bool {
	val, ok := c.Locals("IsAdmin").(bool)
	if !ok {
		return false
	}
	return val
}
