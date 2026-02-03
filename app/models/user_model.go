package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Password     string    `json:"-"`
	TokenVersion int       `gorm:"default:0" json:"token_version"`
	IsAdmin      bool      `gorm:"default:false" json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
}
