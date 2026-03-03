package repo

import (
	"errors"

	"github.com/minhlc98/bookstore/pkg/models"
	"gorm.io/gorm"
)

type AuthorRepo struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) (*AuthorRepo, error) {
	if db == nil {
		return nil, errors.New("repo: db is nil")
	}
	return &AuthorRepo{db: db}, nil
}

func (r *AuthorRepo) Create(author *models.Author) error {
	return r.db.Create(author).Error
}

func (r *AuthorRepo) List() ([]models.Author, error) {
	var authors []models.Author
	if err := r.db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepo) GetByID(id string) (*models.Author, error) {
	var author models.Author
	if err := r.db.Where("id = ?", id).First(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepo) Update(author *models.Author) error {
	return r.db.Save(author).Error
}

func (r *AuthorRepo) Delete(author *models.Author) error {
	return r.db.Delete(author).Error
}

