package repo

import (
	"errors"

	"github.com/minhlc98/bookstore/pkg/models"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) (*BookRepo, error) {
	if db == nil {
		return nil, errors.New("repo: db is nil")
	}
	return &BookRepo{db: db}, nil
}

func (r *BookRepo) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepo) List() ([]models.Book, error) {
	var books []models.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) GetByID(id string) (*models.Book, error) {
	var book models.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepo) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepo) Delete(book *models.Book) error {
	return r.db.Delete(book).Error
}