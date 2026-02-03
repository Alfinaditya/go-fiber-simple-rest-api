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

type GetBooksResponse struct {
	ListResponse[[]BookResponse]
}

type GetBooksWithAuthorResponse struct {
	ListResponse[[]BookWithAuthorResponse]
}
