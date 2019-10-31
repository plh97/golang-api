package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string
	Isbn   string
	Title  string
	Author Author
}

type Author struct {
	FirstName string
	LastName  string
}

var books []Book

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Isbn: "1", Title: "哈里波塔", Author: Author{FirstName: "1", LastName: "laaast"}})
	books = append(books, Book{ID: "2", Isbn: "2", Title: "死侍", Author: Author{FirstName: "1", LastName: "laaast"}})
	books = append(books, Book{ID: "3", Isbn: "3", Title: "超市夜未眠", Author: Author{FirstName: "1", LastName: "laaast"}})
	books = append(books, Book{ID: "4", Isbn: "4", Title: "吉泽明步", Author: Author{FirstName: "1", LastName: "laaast"}})
	r.HandleFunc("/api/book", getBooks).Methods("GET")
	r.HandleFunc("/api/book/create", createBook).Methods("GET")
	r.HandleFunc("/api/book/update/{id}", updateBook).Methods("GET")
	r.HandleFunc("/api/book/delete/{id}", deleteBook).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(srv.ListenAndServe())
}

// 获取所有书
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for _, e := range books {
		if e.ID == params["id"] {
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

// 创建书
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// 更新书
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for i, book := range books {
		if book.ID == params["id"] {
			var book Book
			json.NewDecoder(r.Body).Decode(book)
			books[i] = book
			// books = append(books[:i], book, books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
