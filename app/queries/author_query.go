package queries

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthorQueries struct {
	db *gorm.DB
}

func NewAuthorQueries(db *gorm.DB) *AuthorQueries {
	return &AuthorQueries{db: db}
}

func (q *AuthorQueries) GetAuthors() ([]models.Author, error) {
	var authors []models.Author

	// Query all authors from database
	err := q.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (q *AuthorQueries) GetAuthorsWithBooks() ([]models.Author, error) {
	var authors []models.Author

	// Query all authors and preload books relationship
	err := q.db.Preload("Books").Find(&authors).Error
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (q *AuthorQueries) GetAuthorByID(id uuid.UUID) (*models.Author, error) {
	var author models.Author

	// Query author by ID
	err := q.db.First(&author, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (q *AuthorQueries) GetAuthorByIDWithBooks(id uuid.UUID) (*models.Author, error) {
	var author models.Author

	// Query author by ID and preload books
	err := q.db.Preload("Books").First(&author, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (q *AuthorQueries) CreateAuthor(author *models.Author) error {
	return q.db.Create(author).Error
}

func (q *AuthorQueries) UpdateAuthor(author *models.Author) error {
	return q.db.Save(author).Error
}

func (q *AuthorQueries) DeleteAuthor(id uuid.UUID) error {
	result := q.db.Delete(&models.Author{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
