package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Books struct (model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// get Single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	jsos.NewEncoder(w).Encode(books)
}

// Create new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init Router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book(ID:"1", Isbn:"4644656", Title:"Book one", Author: &Author{
		Firstname:"ABC", Lastname:"XYZ"}))
	books = append(books, Book(ID:"2", Isbn:"4644656", Title:"Book one", Author: &Author{
		Firstname:"ABC", Lastname:"XYZ"}))
	books = append(books, Book(ID:"3", Isbn:"4644656", Title:"Book one", Author: &Author{
		Firstname:"abc", Lastname:"xyz"}))			

	// Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
