package middleware

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/queries"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/platform/database"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func AuthProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "Missing Token",
			})
		}
		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenStr, &utils.Claims{}, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "Invalid Token",
			})
		}
		claims := token.Claims.(*utils.Claims)

		db := database.GetDB()
		authQueries := queries.NewUserQueries(db)
		user, err := authQueries.GetUserByID(claims.UserID)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "User no longer exists",
			})
		}

		if user.TokenVersion != claims.TokenVersion {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "Session revoked. Please log in again.",
			})
		}
		c.Locals("IsAdmin", user.IsAdmin)
		c.Locals("UserID", user.ID.String())
		c.Locals("TokenVersion", user.TokenVersion)
		return c.Next()
	}
}
