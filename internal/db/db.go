package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserS struct {
	Email    string
	Password string
}

func connect() *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
func Insert(email string, password string) (bool, *mongo.InsertOneResult) {
	client := connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("blog").Collection("users")
	doc := bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "password", Value: password}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, result
}
