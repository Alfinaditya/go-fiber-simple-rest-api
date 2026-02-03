package handlers

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/dto"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/models"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/queries"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/pkg/utils"
	"github.com/Alfinaditya/go-fiber-simple-rest-api/platform/database"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get list of all books
// @Tags Books
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.GetBooksResponse  "List of books"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/books [get]
func GetBooks(c *fiber.Ctx) error {
	db := database.GetDB()

	bookQueries := queries.NewBookQueries(db)

	books, err := bookQueries.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse("Internal Server Error"),
		)
	}

	bookResponses := make([]dto.BookResponse, len(books))
	for i, book := range books {
		bookResponses[i] = dto.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			ISBN:        book.ISBN,
			PublishYear: book.PublishYear,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		}
	}
	return c.JSON(dto.NewListResponse(bookResponses, len(bookResponses), ""))
}

// GetBooksWithAuthor godoc
// @Summary Get all books with author details
// @Description Get list of all books including their author information
// @Tags Books
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.GetBooksWithAuthorResponse "List of books with authors"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/books/with-authors [get]
func GetBooksWithAuthor(c *fiber.Ctx) error {
	db := database.GetDB()
	bookQueries := queries.NewBookQueries(db)

	books, err := bookQueries.GetBooksWithAuthors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse("Internal Server Error"),
		)
	}

	bookResponses := make([]dto.BookWithAuthorResponse, len(books))
	for i, book := range books {
		bookResponses[i] = dto.BookWithAuthorResponse{
			ID:          book.ID,
			Title:       book.Title,
			ISBN:        book.ISBN,
			PublishYear: book.PublishYear,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
			AuthorID:    book.AuthorID,
			Author: dto.AuthorResponse{
				ID:        book.Author.ID,
				Name:      book.Author.Name,
				Bio:       book.Author.Bio,
				BirthDate: book.Author.BirthDate,
			},
		}
	}
	return c.JSON(dto.NewListResponse(bookResponses, len(bookResponses), ""))
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided information
// @Tags Books
// @Accept json
// @Produce json
// @Security Bearer
// @Param author body dto.CreateBookRequest true "Book information"
// @Success 201 "Book created successfully"
// @Failure 400 {object} dto.ValidationErrorResponse "Bad request / validation error"
// @Failure 401 {object} dto.BaseResponse "Unauthorized"
// @Failure 500 {object} dto.BaseResponse "Internal server error"
// @Router /api/books [post]
func CreateBook(c *fiber.Ctx) error {
	db := database.GetDB()

	authorQueries := queries.NewBookQueries(db)

	var req dto.CreateBookRequest
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

	author := models.Book{
		ID:          uuid.Must(uuid.NewV7()),
		Title:       req.Title,
		ISBN:        req.ISBN,
		PublishYear: req.PublishYear,
		AuthorID:    req.AuthorId,
	}

	if err := authorQueries.CreateBook(&author); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse("Failed to create book"),
		)
	}

	return c.SendStatus(fiber.StatusCreated)
}
