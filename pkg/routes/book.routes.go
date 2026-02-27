package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhlc98/bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/books/{id}", controllers.UpdateById).Methods("PUT")
}
