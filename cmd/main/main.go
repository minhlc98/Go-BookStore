package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/minhlc98/bookstore/pkg/config"
	"github.com/minhlc98/bookstore/pkg/controllers"
	"github.com/minhlc98/bookstore/pkg/middleware"
	"github.com/minhlc98/bookstore/pkg/models"
	"github.com/minhlc98/bookstore/pkg/repo"
	"github.com/minhlc98/bookstore/pkg/routes"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := models.Migrate(db); err != nil {
		log.Fatal(err)
	}

	var PORT = os.Getenv("PORT")
	var CORS_ORIGIN = os.Getenv("CORS_ORIGIN")
	var APP_ENV = os.Getenv("APP_ENV")

	if PORT == "" {
		log.Panic("PORT is empty")
	}

	if CORS_ORIGIN == "" {
		log.Panic("CORS_ORIGIN is empty")
	}

	if APP_ENV == "" {
		log.Panic("APP_ENV is empty")
	}

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{CORS_ORIGIN},
	})

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

	fmt.Printf("App is running on %s - %s \n", PORT, APP_ENV)
	if err := http.ListenAndServe(PORT, crs.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
