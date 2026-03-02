package models

import (
	"github.com/minhlc98/bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      int    `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) Create() error {
	if err := db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (book *Book) Update() (error) {
	if err := db.Save(book).Error; err != nil {
		return err
	}
	return nil
}

func (book *Book) Delete() error {
	if err := db.Delete(book).Error; err != nil {
		return err
	}
	return nil
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int) (*Book, error) {
	var book Book
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
