package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	ISBN        string    `json:"isbn" gorm:"not null"`
	PublishYear int       `json:"publish_year" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`

	AuthorID uuid.UUID `json:"author_id"`
	Author   Author    `json:"author" gorm:"foreignKey:AuthorID"`
}
