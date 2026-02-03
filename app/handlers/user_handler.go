package handlers

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/queries"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/platform/database"

	"github.com/gofiber/fiber/v2"
)

// Logout godoc
// @Summary Logout user
// @Description Revoke user's token version to invalidate all tokens
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.BaseResponse "Successfully logged out"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Router /api/users/logout [post]
func Logout(c *fiber.Ctx) error {

	db := database.GetDB()

	authQueries := queries.NewAuthQueries(db)

	userId := utils.GetUserID(c)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized",
		})
	}

	authQueries.RevokeTokenVersion(userId)

	return c.SendStatus(fiber.StatusOK)
}

// Me godoc
// @Summary Get current user info
// @Description Get authenticated user's information
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.UserResponse "User information"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 404 {object} dto.BaseResponse "User not found"
// @Router /api/users/me [get]
func Me(c *fiber.Ctx) error {

	db := database.GetDB()

	userQueries := queries.NewUserQueries(db)

	userId := utils.GetUserID(c)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized",
		})
	}

	user, err := userQueries.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Get user successfully",
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"is_admin": user.IsAdmin,
		},
	})
}
