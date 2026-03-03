package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/minhlc98/bookstore/pkg/config"
	"github.com/minhlc98/bookstore/pkg/controllers"
	"github.com/minhlc98/bookstore/pkg/middleware"
	"github.com/minhlc98/bookstore/pkg/models"
	"github.com/minhlc98/bookstore/pkg/repo"
	"github.com/minhlc98/bookstore/pkg/routes"
)

const PORT = ":8080"

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := models.Migrate(db); err != nil {
		log.Fatal(err)
	}

	bookRepo, err := repo.NewBookRepo(db)
	if err != nil {
		log.Fatal(err)
	}
	authorRepo, err := repo.NewAuthorRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(middleware.Recovery)
	routes.RegisterAuthorRoutes(r, controllers.NewAuthorController(authorRepo))
	routes.RegisterBookRoutes(r, controllers.NewBookController(bookRepo))

	fmt.Printf("App is running on %s \n", PORT)
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
