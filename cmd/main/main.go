package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/minhlc98/bookstore/pkg/routes"
)

const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	fmt.Printf("App is running on %s \n", PORT)
	http.ListenAndServe(PORT, r)
}
