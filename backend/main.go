package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// type Trainer struct {
// 	Name string
// 	Age  int
// 	City string
// }

var books []Book
var collection *mongo.Collection
var ctx context.Context

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handleMongodb() {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	collection = client.Database("testing").Collection("numbers")
}

func main() {
	handleMongodb()
	handleRoute()
}

// 获取所有书
func getBooks(w http.ResponseWriter, r *http.Request) {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
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
	// json.NewDecoder(r.Body).Decode(&book)
	// book.ID = strconv.Itoa(rand.Intn(1000000))
	// insertResult, err := collection.InsertOne(context.TODO(), book)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(1111, insertResult)
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

func handleRoute() {
	r := mux.NewRouter()
	books = append(books, Book{ID: strconv.Itoa(rand.Intn(1000000)), Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "哈里波塔"})
	books = append(books, Book{ID: strconv.Itoa(rand.Intn(1000000)), Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "死侍"})
	books = append(books, Book{ID: strconv.Itoa(rand.Intn(1000000)), Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "超市夜未眠"})
	books = append(books, Book{ID: strconv.Itoa(rand.Intn(1000000)), Author: "peng", CreateAt: time.Now(), UpdateAt: time.Now(), Name: "吉泽明步"})
	var interfaceSlice []interface{} = make([]interface{}, len(books))
	for i, d := range books {
		interfaceSlice[i] = d
	}
	collection.InsertMany(context.TODO(), interfaceSlice)
	r.HandleFunc("/api/book/{id}", updateBook).Methods(http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/api/book/{id}", deleteBook).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/book", createBook).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/book", getBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/book/{id}", getBook).Methods(http.MethodGet)
	r.Use(corsMiddleware)
	r.Use(loggingMiddleware)
	http.ListenAndServe("0.0.0.0:8080", r)
}
