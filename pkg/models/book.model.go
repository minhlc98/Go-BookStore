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

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) Book {
	var book Book
	db.Where("id = ?", id).First(&book)
	return book
}

func DeleteBookById(id int64) bool {
	var book Book
	db.Where("id=?", id).Delete(&book)
	return true
}

func UpdateById(id int64, book *Book) *Book {
	db.Save(book)
	return book
}
