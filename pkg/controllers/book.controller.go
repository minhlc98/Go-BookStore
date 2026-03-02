package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/minhlc98/bookstore/pkg/models"
	"github.com/minhlc98/bookstore/pkg/utils"
	"gorm.io/gorm"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	utils.ParseBody(r, &book)
	if err := book.Create(); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	book, err := models.GetBookById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(err)
	}
	json.NewEncoder(w).Encode(book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	book, errGetBook := models.GetBookById(id)
	if errGetBook != nil {
		if errors.Is(errGetBook, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(errGetBook)
	}
	errDeleteBook := book.Delete()
	if errDeleteBook != nil {
		if errors.Is(errDeleteBook, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(errDeleteBook)
	}
	json.NewEncoder(w).Encode(book)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book models.Book
	utils.ParseBody(r, &book)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	bookDetail, err := models.GetBookById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(err)
	}
	bookDetail.Author = book.Author
	bookDetail.Name = book.Name
	bookDetail.Publication = book.Publication

	if errUpdate := bookDetail.Update(); errUpdate != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book Not Found", http.StatusNotFound)
			return
		}
		panic(errUpdate)
	}
	json.NewEncoder(w).Encode(bookDetail)
}
