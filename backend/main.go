package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

// Book is a book title
type Book struct {
	ID       primitive.ObjectID
	Name     string
	Author   string
	CreateAt time.Time
	UpdateAt time.Time
}

var books []Book
var collection *mongo.Collection
var ctx context.Context

// ObjectID id
type ObjectID [12]byte

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
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, _ := collection.Find(ctx, bson.D{})
	res := []bson.M{}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, result)
	}
	json.NewEncoder(w).Encode(res)
}

// 创建书
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	result, err := collection.InsertOne(
		context.Background(),
		bson.D{
			primitive.E{Key: "name", Value: book.Name},
			primitive.E{Key: "author", Value: book.Author},
			primitive.E{Key: "createAt", Value: time.Now()},
			primitive.E{Key: "updateAt", Value: time.Now()},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

// 更新书
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	objectID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	result, err := collection.UpdateOne(
		context.Background(),
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
		},
		bson.D{
			primitive.E{Key:"$set",Value: bson.D{
				primitive.E{Key: "name", Value: book.Name},
				primitive.E{Key: "author", Value: book.Author},
				primitive.E{Key: "updateAt", Value: time.Now()},
			}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params
	objectID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	result, err := collection.DeleteOne(
		context.Background(),
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bson.M{
		"name": "book.Name",
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bson.M{
		"name": "book.Name",
		"token": "wefrtykujhgf234567u",
	})
}

func handleRoute() {
	r := mux.NewRouter()
	r.HandleFunc("/api/register", handleRegister).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/login", handleLogin).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/book", createBook).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/book/{id}", updateBook).Methods(http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/api/book/{id}", deleteBook).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/book", getBooks).Methods(http.MethodGet)
	r.Use(corsMiddleware)
	r.Use(loggingMiddleware)
	http.ListenAndServe(":8080", r)
}
