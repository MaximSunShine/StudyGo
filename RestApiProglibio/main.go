package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Author *Authors `json:"author"`
}
type Authors struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	r := mux.NewRouter()
	books = append(books, Book{
		ID:    "1",
		Title: "Война и Мир",
		Author: &Authors{
			Firstname: "Лев",
			Lastname:  "Толстой",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Title: "Преступление и наказание",
		Author: &Authors{
			Firstname: "Фёдор",
			Lastname:  "Достоевский",
		},
	})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
