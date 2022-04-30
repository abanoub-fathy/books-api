package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize books
	InitializeBooks()

	// declare a router handler
	r := mux.NewRouter()

	// create endpoints
	r.HandleFunc("/api/v1/books", CreateBook).Methods("POST")
	r.HandleFunc("/api/v1/books", GetAllBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", GetSingleBook).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", UpdateSingleBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/{id}", DeleteBook).Methods("DELETE")

	// launch server
	fmt.Println("App is working on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
