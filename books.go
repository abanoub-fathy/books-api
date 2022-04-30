package main

import "errors"

var books []Book

// Initialize the books data
func InitializeBooks() {
	book1 := Book{Id: "1", Isbn: "212", Title: "book1", Pages: 20, Author: Author{Id: "1", FirstName: "sd1", LastName: "sds1"}}
	book2 := Book{Id: "2", Isbn: "212", Title: "book2", Pages: 40, Author: Author{Id: "2", FirstName: "sd2", LastName: "sds2"}}

	// append book1 to books
	books = append(books, book1, book2)
}

// Get Books From the DB
func GetBooksFromDB() []Book {
	return books
}

// Get single book by id
func GetSingleBookById(id string) (book Book, err error) {
	// loop through all books
	for _, book := range books {
		if book.Id == id {
			return book, nil
		}
	}

	return Book{}, errors.New("not found any book")
}
