package main

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	apartment  Apartment
	apartments []Apartment
)

func TestPostApartmentEndpoint(t *testing.T) {
	apartment.Name = "Ascott"
	apartment.Location = "Kuningan"
	apartment.Price = 150000000
	apartment.Availabilty = false
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	if err != nil {
		t.Errorf("Terjadi error ketika koneksi ke client server")
	}
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	insertResult, err := apartmentCollection.InsertOne(ctx, apartment)
	if err != nil {
		t.Errorf("GAGAL! Terjadi error ketika menginput data ke collection")
	}
	if insertResult.InsertedID == nil {
		t.Errorf("GAGAL! Terjadi error ketika menginput data ke collection")
	}
}

func TestGetApartmentsEndpoint(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	result, err := apartmentCollection.Find(ctx, bson.M{})
	if err != nil {
		t.Errorf("Terjadi error ketika menarik data dari koleksi!")
	}
	for result.Next(ctx) {
		var apartment Apartment
		result.Decode(&apartment)
		apartments = append(apartments, apartment)
	}
	if err := result.Err(); err != nil {
		t.Errorf("Terjadi error ketika menarik data dari koleksi!")
	}
}

func TestUpdateApartmentEndpoint(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	if err != nil {
		t.Errorf("Terjadi error saat mencoba konek ke server!")
	}
	update := bson.D{
		{"$inc", bson.D{
			{"price", 125000000},
		}},
	}
	updateResult, err := apartmentCollection.UpdateOne(ctx, bson.D{}, update)
	if err != nil {
		t.Errorf("Terjadi error saat mencoba update data!")
	}
	if updateResult.MatchedCount == 0 {
		t.Errorf("GAGAL!! Tidak berhasil mengupdate apapun dari koleksi tersebut!")
	}
}

func TestDeleteApartmentEndpoint(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://spacestock:qwerty123@cluster1-shard-00-00-n2kum.mongodb.net:27017,cluster1-shard-00-01-n2kum.mongodb.net:27017,cluster1-shard-00-02-n2kum.mongodb.net:27017/test?ssl=true&replicaSet=cluster1-shard-0&authSource=admin&retryWrites=true&w=majority"))
	if err != nil {
		t.Errorf("Terjadi error saat mencoba konek ke server!")
	}
	apartmentCollection := client.Database("spacesctock").Collection("apartments")
	result, err := apartmentCollection.DeleteOne(ctx, bson.D{})
	if err != nil {
		t.Errorf("Terjadi error saat mencoba menghapus koleksi!")
	}
	if result.DeletedCount == 0 {
		t.Errorf("GAGAL!! Tidak menghapus apapun!")
	}
}
