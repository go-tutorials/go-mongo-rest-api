package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id,omitempty" firestore:"id,omitempty" validate:"required,max=40"`
	Username    string     `json:"username,omitempty" gorm:"column:username" bson:"username,omitempty" dynamodbav:"username,omitempty" firestore:"username,omitempty" validate:"required,username,max=100"`
	Email       string     `json:"email,omitempty" gorm:"column:email" bson:"email,omitempty" dynamodbav:"email,omitempty" firestore:"email,omitempty" validate:"email,max=100"`
	Phone       string     `json:"phone,omitempty" gorm:"column:phone" bson:"phone,omitempty" dynamodbav:"phone,omitempty" firestore:"required,phone,omitempty" validate:"required,phone,max=18"`
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty" gorm:"column:date_of_birth" bson:"dateOfBirth,omitempty" dynamodbav:"dateOfBirth,omitempty" firestore:"dateOfBirth,omitempty"`
}

type Handler struct {
	Collection *mongo.Collection
}

func (db *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panicf("%s : %s", "error while get cursor users", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var users []User
	for cursor.Next(context.TODO()) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			log.Panicf("%s : %s", "error while get user from curor", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(users)
	if err != nil {
		log.Panicf("%s : %s", "error while parse json to reponse", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (db *Handler) Load(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	query := bson.M{"_id": id}
	result := db.Collection.FindOne(context.TODO(), query)
	if result.Err() != nil {
		if strings.Compare(fmt.Sprint(result.Err()), "mongo: no documents in result") == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("null"))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(result.Err().Error()))
			return
		}
	}
	user := User{}
	err := result.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(user)
	if err != nil {
		log.Panicf("%s : %s", "error while parse json to respone", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (db *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	postBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panicf("%s : %s", "error while read post body", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(postBody, &user)
	if err != nil {
		log.Panicf("%s : %s", "error while parse to struct user", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := db.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(result)
	if err != nil {
		log.Panicf("%s : %s", "error while parse to json to respone", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (db *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	var user User
	putBody, er0 := ioutil.ReadAll(r.Body)
	if er0 != nil {
		log.Panicf("%s : %s", "error while read put body", er0)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er0.Error()))
		return
	}

	er1 := json.Unmarshal(putBody, &user)
	if er1 != nil {
		log.Panicf("%s : %s", "error while parse to struct user", er1)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er1.Error()))
		return
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	result, er2 := db.Collection.UpdateOne(context.TODO(), filter, update)
	if er2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er2.Error()))
		return
	}
	response, er3 := json.Marshal(result)
	if er3 != nil {
		log.Panicf("%s : %s", "error while parse json to respone", er3)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er3.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (db *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}
	result, er0 := db.Collection.DeleteOne(context.TODO(), filter)
	if er0 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er0.Error()))
		return
	}
	response, er1 := json.Marshal(result)
	if er1 != nil {
		log.Panicf("%s : %s", "error while parse json to respone", er1)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er1.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	fmt.Println("Connecting to mongo...")
	uri := "mongodb://localhost:27017"
	database := "master_data"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongo")
	userCollection := client.Database(database).Collection("users")
	db := Handler{Collection: userCollection}

	r := mux.NewRouter()
	r.HandleFunc("/users", db.Create).Methods("POST")
	r.HandleFunc("/users/{id}", db.Load).Methods("GET")
	r.HandleFunc("/users", db.GetAll).Methods("GET")
	r.HandleFunc("/users/{id}", db.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", db.Delete).Methods("DELETE")

	fmt.Println("Start server at port :8080")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatalln(srv.ListenAndServe())
}
