package dto

import (
	"time"

	"github.com/google/uuid"
)

type BookResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ISBN        string    `json:"isbn"`
	PublishYear int       `json:"publish_year"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookWithAuthorResponse struct {
	ID          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	ISBN        string         `json:"isbn"`
	PublishYear int            `json:"publish_year"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	AuthorID    uuid.UUID      `json:"author_id"`
	Author      AuthorResponse `json:"author"`
}

type CreateBookRequest struct {
	Title       string    `json:"title" validate:"required,max=100"`
	ISBN        string    `json:"isbn" validate:"required,max=500"`
	PublishYear int       `json:"publish_year" validate:"required"`
	AuthorId    uuid.UUID `json:"author_id" validate:"required"`
}

type GetBooksResponse struct {
	ListResponse[[]BookResponse]
}

type GetBooksWithAuthorResponse struct {
	ListResponse[[]BookWithAuthorResponse]
}
