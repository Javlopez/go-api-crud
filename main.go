package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"fmt"
	"log"
	"wiki/app/driver"
	"wiki/app/handler/http"
	"github.com/go-chi/chi/middleware"
)

const (
	PORT string = ":5000"
)

func main() {

	dbName := "go-mysql-crud"
	dbPass := "root"
	dbHost := "localhost"
	dbPort := "3306"

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	postHandler := handler.NewPostHandler(connection)
	r.Get("/posts", postHandler.Fetch)
	r.Get("/posts/{id}", postHandler.GetByID)
	r.Post("/posts/create", postHandler.Create)
	r.Put("/posts/update/{id}", postHandler.Update)
	r.Delete("/posts/{id}", postHandler.Delete)

	fmt.Println("Server listen at 5000")
	err = http.ListenAndServe(PORT, r)

	if err != nil {
		log.Fatal(err)
	}
}

