package main

// append data to the books
func GetBooksFromDB() []Book {
	// slice of books
	books := []Book{}
	book1 := Book{Id: "1", Isbn: "212", Title: "book1", Pages: 20, Author: Author{Id: "1", FirstName: "sd1", LastName: "sds1"}}
	book2 := Book{Id: "2", Isbn: "212", Title: "book2", Pages: 40, Author: Author{Id: "2", FirstName: "sd2", LastName: "sds2"}}

	// append book1 to books
	books = append(books, book1, book2)

	return books
}
