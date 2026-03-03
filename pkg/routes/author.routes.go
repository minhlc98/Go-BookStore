package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhlc98/bookstore/pkg/controllers"
)

func RegisterAuthorRoutes(router *mux.Router, authorController *controllers.AuthorController) {
	router.HandleFunc("/authors", authorController.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors", authorController.GetAllAuthor).Methods("GET")
	router.HandleFunc("/authors/{id}", authorController.GetAuthor).Methods("GET")
	router.HandleFunc("/authors/{id}", authorController.DeleteAuthorById).Methods("DELETE")
	router.HandleFunc("/authors/{id}", authorController.UpdateAuthorById).Methods("PUT")
}
