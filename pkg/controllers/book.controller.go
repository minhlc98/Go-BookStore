package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/minhlc98/bookstore/pkg/models"
	"github.com/minhlc98/bookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	utils.ParseBody(r, &book)
	newBook, err := book.CreateBook()
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	book := models.GetBookById(id)
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
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	isSuccess := models.DeleteBookById(id)
	res := map[string]bool{
		"success": isSuccess,
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book models.Book
	utils.ParseBody(r, &book)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		panic(err)
	}
	bookDetail := models.GetBookById(id)
	bookDetail.Author = book.Author
	bookDetail.Name = book.Name
	bookDetail.Publication = book.Publication

	models.UpdateById(id, &bookDetail)
	json.NewEncoder(w).Encode(bookDetail)
}
