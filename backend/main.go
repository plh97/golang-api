package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book is a book struct
type Book struct {
	id     			string
	name  			string
	author 			string
	createTime 	int
	updateTime 	int
}

var books []Book

func main() {
	r := mux.NewRouter()
	books = append(books, Book{id: "1", createTime: 2016, updateTime: 2017 , author: "Bob", name: "哈里波塔"})
	books = append(books, Book{id: "2", createTime: 2016, updateTime: 2017 , author: "Bob", name: "死侍"})
	books = append(books, Book{id: "3", createTime: 2016, updateTime: 2017 , author: "Bob", name: "超市夜未眠"})
	books = append(books, Book{id: "4", createTime: 2016, updateTime: 2017 , author: "Bob", name: "吉泽明步"})
	r.HandleFunc("/api/book", getBooks).Methods("GET")
	r.HandleFunc("/api/book/create", createBook).Methods("GET")
	r.HandleFunc("/api/book/update/{id}", updateBook).Methods("GET")
	r.HandleFunc("/api/book/delete/{id}", deleteBook).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
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
		if e.id == params["id"] {
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
	book.id = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// 更新书
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for i, book := range books {
		if book.id == params["id"] {
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
		if book.id == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
