package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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
func generateHash(pw string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), 10)
	return string(hash)
}

func RegisterUser(email string, password string) bool {
	client := connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	pwHash := generateHash(password)

	coll := client.Database("blog").Collection("users")
	doc := bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "password", Value: pwHash}}

	_, err := coll.InsertOne(context.TODO(), doc)

	return err != nil
}
func LoginUser(email string, password string) bool {
	client := connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var dbUser UserS
	coll := client.Database("blog").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := coll.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)
	if err != nil {
		log.Println(err)
		return false
	}
	pwHash := dbUser.Password
	err = bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(password))

	return err != nil
}
