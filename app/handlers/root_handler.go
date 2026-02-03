package handlers

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/dto"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/queries"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RevokeTokenVersion godoc
// @Summary Revoke user token version (Admin only)
// @Description Revoke a user's token version to invalidate all their tokens. Admin only.
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "User ID"
// @Success 200 "Token version revoked"
// @Failure 400 {object} dto.BaseResponse "Bad request"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Router /api/root/revoke/{id} [post]
func RevokeTokenVersion(c *fiber.Ctx) error {
	db := database.GetDB()

	authQueries := queries.NewAuthQueries(db)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID format",
		})
	}

	userId := utils.GetUserID(c)

	if userId != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized",
		})
	}

	if userId == id.String() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "You cannot revoke your own token",
		})

	}

	authQueries.RevokeTokenVersion(id.String())

	return c.SendStatus(fiber.StatusOK)
}

// GetUsers godoc
// @Summary Get all users (Admin only)
// @Description Get list of all users. Admin only.
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.GetUsersResponse "List of users"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/root/users [get]
func GetUsers(c *fiber.Ctx) error {
	db := database.GetDB()

	userQueries := queries.NewUserQueries(db)

	users, err := userQueries.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse("Internal Server Error"))
	}

	return c.JSON(dto.NewListResponse(users, len(users), ""))
}
