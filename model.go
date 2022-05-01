package main

// model of book
type Book struct {
	Id     string `json:"id"`
	Isbn   string `json:"isbn"`
	Pages  int64  `json:"pages"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// model of author
type Author struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
