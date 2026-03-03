package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhlc98/bookstore/pkg/controllers"
)

func RegisterBookRoutes(router *mux.Router, bookController *controllers.BookController) {
	router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/books", bookController.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/books/{id}", bookController.UpdateById).Methods("PUT")
}
