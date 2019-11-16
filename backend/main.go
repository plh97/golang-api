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

// Account is a account type
type Account struct {
	Name     string
	Password string
}

var books []Book
var booksCollection *mongo.Collection
var accountsCollection *mongo.Collection
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
	booksCollection = client.Database("testing").Collection("books")
	accountsCollection = client.Database("testing").Collection("accounts")
}

func main() {
	handleMongodb()
	handleRoute()
}

// 获取所有书
func getBooks(w http.ResponseWriter, r *http.Request) {
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, _ := booksCollection.Find(ctx, bson.D{})
	res := []bson.M{}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, result)
	}
	json.NewEncoder(w).Encode(
		bson.M{
			"errorCode": 0,
			"data":      res,
		},
	)
}

// 创建书
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	result, err := booksCollection.InsertOne(
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
	json.NewEncoder(w).Encode(
		bson.M{
			"errorCode": 0,
			"data":      result,
		},
	)
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
	result, err := booksCollection.UpdateOne(
		context.Background(),
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
		},
		bson.D{
			primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "name", Value: book.Name},
				primitive.E{Key: "author", Value: book.Author},
				primitive.E{Key: "updateAt", Value: time.Now()},
			}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(
		bson.M{
			"errorCode": 0,
			"data":      result,
		},
	)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params
	objectID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	result, err := booksCollection.DeleteOne(
		context.Background(),
		bson.D{
			primitive.E{Key: "_id", Value: objectID},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(
		bson.M{
			"errorCode": 0,
			"data":      result,
		},
	)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var account Account
	var resultFind struct {
		Value float64
	}
	json.NewDecoder(r.Body).Decode(&account)
	errFind := accountsCollection.FindOne(
		context.Background(),
		bson.M{
			"name":     account.Name,
			"password": account.Password,
		},
	).Decode(&resultFind)
	if errFind != nil {
		json.NewEncoder(w).Encode(bson.M{
			"message":   "account name or password not right",
			"errorCode": 1,
		})
		return
	}
	json.NewEncoder(w).Encode(bson.M{
		"errorCode": 0,
		"data": bson.M{
			"name":  "book.Name",
			"token": "2345uilerghtjyukr",
		},
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var account Account
	var resultFind struct {
		Value float64
	}
	json.NewDecoder(r.Body).Decode(&account)
	errFind := accountsCollection.FindOne(
		context.Background(),
		bson.M{"name": account.Name},
	).Decode(&resultFind)
	if errFind == nil {
		json.NewEncoder(w).Encode(bson.M{
			"message":   "account has been register",
			"errorCode": 1,
		})
		return
	}
	accountsCollection.InsertOne(
		context.Background(),
		bson.D{
			primitive.E{Key: "name", Value: account.Name},
			primitive.E{Key: "password", Value: account.Password},
		},
	)
	json.NewEncoder(w).Encode(bson.M{
		"errorCode": 0,
		"data": bson.M{
			"name":  "book.Name",
			"token": "2345uilerghtjyukr",
		},
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
