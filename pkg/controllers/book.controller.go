package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhlc98/bookstore/pkg/models"
	"github.com/minhlc98/bookstore/pkg/repo"
	"github.com/minhlc98/bookstore/pkg/utils"
	"gorm.io/gorm"
)

type BookController struct {
	repo *repo.BookRepo
}

func NewBookController(r *repo.BookRepo) *BookController {
	return &BookController{repo: r}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	if err := c.repo.Create(&book); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := c.repo.GetByID(vars["id"])
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(err)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.repo.List()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (c *BookController) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, errGetBook := c.repo.GetByID(vars["id"])
	if errGetBook != nil {
		if errors.Is(errGetBook, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(errGetBook)
	}
	errDeleteBook := c.repo.Delete(book)
	if errDeleteBook != nil {
		if errors.Is(errDeleteBook, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(errDeleteBook)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) UpdateById(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	vars := mux.Vars(r)
	bookDetail, err := c.repo.GetByID(vars["id"])
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(err)
	}
	bookDetail.AuthorId = book.AuthorId
	bookDetail.Name = book.Name
	bookDetail.Publication = book.Publication

	if errUpdate := c.repo.Update(bookDetail); errUpdate != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(errUpdate)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(bookDetail)
}
