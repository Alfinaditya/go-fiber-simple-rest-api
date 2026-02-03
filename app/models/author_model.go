package models

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Bio       string    `json:"bio" gorm:"not null"`
	BirthDate string    `json:"birth_date" gorm:"type:varchar(10);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	Books []Book `json:"books,omitempty" gorm:"foreignKey:AuthorID"`
}
