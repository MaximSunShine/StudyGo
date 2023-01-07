package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

// получаем весь список книг
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// получаем книгу по айди параметру передаваемый в урл запросе
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _, v := range books {
		if v.ID == p["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Создаем новую книгу
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Обновляем данные по книге по ее айди
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	var upBook Book
	for i, v := range books {
		if v.ID == p["id"] {
			_ = json.NewDecoder(r.Body).Decode(&upBook)
			books[i].Title = upBook.Title
			books[i].Author = upBook.Author
			break
		}
	}
	json.NewEncoder(w).Encode(upBook)
}

// Удаляем книгу по айди
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for i, v := range books {
		if v.ID == p["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
