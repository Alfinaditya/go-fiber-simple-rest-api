package queries

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/models"

	"gorm.io/gorm"
)

type UserQueries struct {
	db *gorm.DB
}

func NewUserQueries(db *gorm.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (q *UserQueries) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := q.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (q *UserQueries) GetUserByID(id string) (*models.User, error) {
	var user models.User

	err := q.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	var users []models.User

	err := q.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
