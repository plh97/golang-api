package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Book is a book title
type Book struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

var books []Book

func main() {
	r := mux.NewRouter()

	// r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	// 	w.WriteHeader(http.StatusNoContent)
	// 	return
	// })
	books = append(books, Book{ID: "1", Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "哈里波塔"})
	books = append(books, Book{ID: "2", Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "死侍"})
	books = append(books, Book{ID: "3", Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "超市夜未眠"})
	books = append(books, Book{ID: "4", Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "吉泽明步"})
	// r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/api/book", getBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/book/", createBook).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/book/{id}", updateBook).Methods(http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/api/book/{id}", deleteBook).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/book/{id}", getBook).Methods(http.MethodGet)
	r.Use(loggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8080", r)
	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "0.0.0.0:8080",
	// }
	// log.Fatal(srv.ListenAndServe())
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
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// 更新书
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params
	for i, book := range books {
		if book.ID == params["id"] {
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			books[i].Name = book.Name
			books[i].Author = book.Author
			books[i].UpdateAt = time.Now()
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
