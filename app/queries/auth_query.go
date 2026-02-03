package queries

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/models"

	"gorm.io/gorm"
)

type AuthQueries struct {
	db *gorm.DB
}

func NewAuthQueries(db *gorm.DB) *AuthQueries {
	return &AuthQueries{db: db}
}

func (q *AuthQueries) RevokeTokenVersion(userID string) error {
	return q.db.Model(&models.User{}).Where("id = ?", userID).
		Update("token_version", gorm.Expr("token_version + 1")).Error
}
