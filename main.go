package main

import (
	"github.com/EmreSahna/go_mysql_book_management/middleware"
	routes "github.com/EmreSahna/go_mysql_book_management/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	routes.RegisterBookStoreRoutes(r)
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
