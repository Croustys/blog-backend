package db

import (
	"blog-backend/internal/types"
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

func RegisterUser(email string, password string, username string) bool {
	client := connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	pwHash := generateHash(password)

	coll := client.Database("blog").Collection("users")
	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Println(err)
	}

	doc := bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "username", Value: username}, primitive.E{Key: "password", Value: pwHash}}

	_, err = coll.InsertOne(context.TODO(), doc)

	return err == nil
}
func LoginUser(email string, password string) (bool, string) {
	client := connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var dbUser types.UserS
	coll := client.Database("blog").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := coll.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)
	if err != nil {
		log.Println(err)
		return false, ""
	}
	pwHash := dbUser.Password
	err = bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(password))

	return err == nil, dbUser.Username
}

func SavePost(authorEmail string, authorUsername string, title string, content string) bool {
	client := connect()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("blog").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.InsertOne(ctx, bson.D{primitive.E{Key: "author", Value: authorEmail}, primitive.E{Key: "username", Value: authorUsername}, primitive.E{Key: "title", Value: title}, primitive.E{Key: "content", Value: content}})
	return err == nil
}
func GetPosts(offset int64, limit int64) []types.PostResponse {
	client := connect()

	opts := options.Find()
	opts.SetSkip(offset)
	opts.SetLimit(limit)

	coll := client.Database("blog").Collection("posts")
	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println(err)
	}

	var results []types.PostResponse
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
	}

	return results
}
func GetPost(id string) types.PostResponse {
	client := connect()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	coll := client.Database("blog").Collection("posts")
	cursor := coll.FindOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}

	var result types.PostResponse
	cursor.Decode(&result)

	return result
}
