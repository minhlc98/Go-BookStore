package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/minhlc98/bookstore/pkg/middleware"
	"github.com/minhlc98/bookstore/pkg/routes"
)

const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.Use(middleware.Recovery)
	routes.RegisterBookStoreRoutes(r)

	fmt.Printf("App is running on %s \n", PORT)
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
