package handlers

import (
	"errors"
	"strings"

	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/dto"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/models"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/queries"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetAuthors godoc
// @Summary Get all authors
// @Description Get list of all authors
// @Tags Authors
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200  {object} dto.GetAuthorsResponse "List of authors"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/authors [get]
func GetAuthors(c *fiber.Ctx) error {
	db := database.GetDB()

	authorQueries := queries.NewAuthorQueries(db)

	authors, err := authorQueries.GetAuthors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse("Internal Server Error"),
		)
	}

	return c.JSON(dto.NewListResponse(authors, len(authors), ""))
}

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the provided information
// @Tags Authors
// @Accept json
// @Produce json
// @Security Bearer
// @Param author body dto.CreateAuthorRequest true "Author information"
// @Success 201 "Author created successfully"
// @Failure 400 {object} dto.ValidationErrorResponse "Bad request / validation error"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/authors [post]
func CreateAuthor(c *fiber.Ctx) error {
	db := database.GetDB()

	authorQueries := queries.NewAuthorQueries(db)

	var req dto.CreateAuthorRequest
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

	author := models.Author{
		ID:        uuid.Must(uuid.NewV7()),
		Name:      req.Name,
		Bio:       req.Bio,
		BirthDate: req.BirthDate,
	}

	if err := authorQueries.CreateAuthor(&author); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse("Failed to create author"),
		)
	}

	return c.SendStatus(fiber.StatusCreated)

}

// UpdateAuthor godoc
// @Summary Update an author
// @Description Update an existing author's information
// @Tags Authors
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Author ID"
// @Param author body dto.UpdateAuthorRequest true "Updated author information"
// @Success 200 {object} dto.DataResponse[models.Author] "Author updated successfully"
// @Failure 400 {object} dto.ValidationErrorResponse "Bad request / validation error"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 404 {object} dto.BaseResponse "Author not found"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/authors/{id} [put]
func UpdateAuthor(c *fiber.Ctx) error {
	db := database.GetDB()
	authorQueries := queries.NewAuthorQueries(db)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse("Invalid author ID format"),
		)
	}

	var req dto.UpdateAuthorRequest
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

	author, err := authorQueries.GetAuthorByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			dto.ErrorResponse("Author not found"),
		)
	}

	utils.UpdateIfNotNil(&author.Name, req.Name)
	utils.UpdateIfNotNil(&author.Bio, req.Bio)
	utils.UpdateIfNotNil(&author.BirthDate, req.BirthDate)

	if err := authorQueries.UpdateAuthor(author); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse("Failed to update author"),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.NewDataResponse(author, "Author updated successfully"),
	)
}

// DeleteAuthor godoc
// @Summary Delete an author
// @Description Delete an author by ID
// @Tags Authors
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Author ID"
// @Success 200 {object} dto.BaseResponse "Author deleted successfully"
// @Failure 400 {object} dto.BaseResponse "Invalid ID format"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 404 {object} dto.BaseResponse "Author not found"
// @Failure 409 {object} dto.BaseResponse "Cannot delete author with existing books"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/authors/{id} [delete]
func DeleteAuthor(c *fiber.Ctx) error {
	db := database.GetDB()
	authorQueries := queries.NewAuthorQueries(db)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse("Invalid author ID format"),
		)
	}

	if err := authorQueries.DeleteAuthor(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				dto.ErrorResponse("Author not found"),
			)
		}
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return c.Status(fiber.StatusConflict).JSON(
				dto.ErrorResponse("Cannot delete author with existing books"),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse("Failed to delete author"),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.SuccessResponse("Author deleted successfully"),
	)
}
