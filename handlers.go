package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// create book handler
func CreateBook(w http.ResponseWriter, r *http.Request) {

}

// get all books handler
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := GetBooksFromDB()
	json.NewEncoder(w).Encode(books)
}

// get single book handler
func GetSingleBook(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// get the id of the required book
	id := mux.Vars(r)["id"]

	// get single book by id
	book, err := GetSingleBookById(id)

	if err != nil {
		// not found
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// return the book
	json.NewEncoder(w).Encode(book)
}

// update a book handler
func UpdateSingleBook(w http.ResponseWriter, r *http.Request) {

}

// delete a book handler
func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
