package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Apartment struct
type Apartment struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Location    string             `json:"location,omitempty" bson:"location,omitempty"`
	Price       int                `json:"price,omitempty" bson:"price,omitempty"`
	Availabilty bool               `json:"availability,omitempty" bson:"availabilty,omitempty"`
}

//CreateApartmentEndpoint {POST}
func CreateApartmentEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var a Apartment
	json.NewDecoder(r.Body).Decode(&a)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	result, err := apartmentCollection.InsertOne(ctx, a)
	if err != nil {
		panic(err.Error())
	}
	r.Body.Close()
	json.NewEncoder(w).Encode(result)
}

// GetApartmentsEndpoint {GET}
func GetApartmentsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var apartments []Apartment
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	result, err := apartmentCollection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	for result.Next(ctx) {
		var apartment Apartment
		result.Decode(&apartment)
		apartments = append(apartments, apartment)
	}
	if err := result.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(apartments)
}

//DeleteApartmentEndpoint {DELETE}
func DeleteApartmentEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	result, err := apartmentCollection.DeleteOne(ctx, Apartment{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

//UpdateApartmentEndpoint {PUT}
func UpdateApartmentEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	update := bson.D{
		{"$inc", bson.D{
			{"price", 125000000},
		}},
	}
	updateResult, err := apartmentCollection.UpdateOne(ctx, Apartment{ID: id}, update)
	json.NewEncoder(w).Encode(updateResult)
}

func main() {
	fmt.Println("Application Starting up at http://localhost:8000")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err.Error())
	}
	r := mux.NewRouter()
	r.HandleFunc("/apartment", CreateApartmentEndpoint).Methods("POST")
	r.HandleFunc("/apartment", GetApartmentsEndpoint).Methods("GET")
	r.HandleFunc("/apartment/{id}", DeleteApartmentEndpoint).Methods("DELETE")
	r.HandleFunc("/apartment/{id}", UpdateApartmentEndpoint).Methods("PUT")
	http.ListenAndServe(":8000", r)
}
