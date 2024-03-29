package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"strconv"

	//"math/rand"
	"net/http"
	//"strconv"
)

// Book Model

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = r.ParseForm()
	var book Book = Book{
		ID:    strconv.Itoa(rand.Intn(100000000)),
		Isbn:  r.FormValue("isbn"),
		Title: r.FormValue("title"),
		Author: Author{
			Firstname: r.FormValue("author[firstname]"),
			Lastname:  r.FormValue("author[lastname]"),
		},
	}
	books = append(books, book)
	_ = json.NewEncoder(w).Encode(book)

}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	r := mux.NewRouter()

	books = append(books, Book{"1", "4489009", "Book one", Author{"John", "Doe"}})
	books = append(books, Book{"2", "45486", "Book two", Author{"Steve", "Smith"}})
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
