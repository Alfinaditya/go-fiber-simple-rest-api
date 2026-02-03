package queries

import (
	"github.com/Alfinaditya/go-fiber-simple-rest-api/app/models"

	"gorm.io/gorm"
)

type BookQueries struct {
	db *gorm.DB
}

func NewBookQueries(db *gorm.DB) *BookQueries {
	return &BookQueries{db: db}
}

func (q *BookQueries) GetBooks() ([]models.Book, error) {
	var books []models.Book

	err := q.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (q *BookQueries) GetBooksWithAuthors() ([]models.Book, error) {
	var books []models.Book

	err := q.db.Preload("Author").Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (q *BookQueries) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book

	err := q.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (q *BookQueries) GetBookByIDWithAuthor(id uint) (*models.Book, error) {
	var book models.Book

	err := q.db.Preload("Author").First(&book, id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (q *BookQueries) CreateBook(book *models.Book) error {
	return q.db.Create(book).Error
}

func (q *BookQueries) UpdateBook(book *models.Book) error {
	return q.db.Save(book).Error
}

func (q *BookQueries) DeleteBook(id uint) error {
	return q.db.Delete(&models.Book{}, id).Error
}
