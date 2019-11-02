package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Todo is a todo list struct
type Todo struct {
	id        string    `json:id`
	name      string    `json:"name"`
	completed bool      `json:"completed"`
	due       time.Time `json:"due"`
}

var todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	todos = []Todo{}
	todos = append(todos, Todo{id: "1", name: "peng"})
	todos = append(todos, Todo{id: "2", name: "zhang"})
	todos = append(todos, Todo{id: "3", name: "dai"})
	todos = append(todos, Todo{id: "4", name: "lin"})
	router.HandleFunc("/api/book", getBooks).Methods("GET")
	router.HandleFunc("/api/book/create", createBook).Methods("GET")
	router.HandleFunc("/api/book/update/{id}", updateBook).Methods("GET")
	router.HandleFunc("/api/book/delete/{id}", deleteBook).Methods("GET")
	router.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(srv.ListenAndServe())
}

// 获取所有书
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(todos)
	json.NewEncoder(w).Encode(todos)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for _, e := range todos {
		if e.id == params["id"] {
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	json.NewEncoder(w).Encode(Todo{})
}

// 创建书
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.id = strconv.Itoa(rand.Intn(10000))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todos)
}

// 更新书
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for i, todo := range todos {
		if todo.id == params["id"] {
			var book Todo
			json.NewDecoder(r.Body).Decode(&book)
			todos[i] = book
			// books = append(books[:i], book, books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for i, todo := range todos {
		if todo.id == params["id"] {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
