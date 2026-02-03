package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"-"`
	TokenVersion int       `json:"token_version"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
}
