package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateAuthorRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=100"`
	Bio       string `json:"bio" validate:"required,min=10,max=500"`
	BirthDate string `json:"birth_date" validate:"required,datetime=2006-01-02"`
}

type UpdateAuthorRequest struct {
	Name      *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Bio       *string `json:"bio,omitempty" validate:"omitempty,min=10,max=500"`
	BirthDate *string `json:"birth_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

type AuthorResponse struct {
	ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name      string    `json:"name" example:"J.K. Rowling"`
	Bio       string    `json:"bio" example:"British author, best known for Harry Potter"`
	BirthDate string    `json:"birth_date" example:"1965-07-31"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-15T10:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-15T10:00:00Z"`
}

type GetAuthorsResponse struct {
	ListResponse[[]AuthorResponse]
}
