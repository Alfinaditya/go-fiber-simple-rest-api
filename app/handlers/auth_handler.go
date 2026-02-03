package handlers

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/dto"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/queries"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/platform/database"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login user
// @Description Authenticate user using username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginAuthDto true "Login payload"
// @Success 200 {object} dto.LoginResponse "token"
// @Failure 400 {object} dto.ValidationErrorResponse "Validation error"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Router /api/auth/login [post]
func Login(c *fiber.Ctx) error {
	db := database.GetDB()

	userQueries := queries.NewUserQueries(db)

	var req dto.LoginAuthDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse("Cannot parse JSON request body"),
		)
	}

	if errors := utils.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ValidationErrorResponseFunc(errors),
		)
	}
	user, err := userQueries.GetUserByUsername(req.Username)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse("Login failed"))
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse("Login failed"))
	}

	token, err := utils.GenerateToken(user.ID.String(), user.TokenVersion, user.IsAdmin)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse("Token generation failed"))
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
