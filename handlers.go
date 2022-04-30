package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// create book handler
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// define book and decode the body
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	// check book validity
	if book.Title == "" || book.Pages == 0 || book.Isbn == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// add the book to the DB
	addedBook, err := AddBookToDB(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedBook)
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
