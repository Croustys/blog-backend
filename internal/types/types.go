package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserS struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type PostS struct {
	AuthorEmail string `json:"authorEmail"`
	Username    string `json:"username"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
type PostResponse struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `json:"username"`
	Title    string             `json:"title"`
	Content  string             `json:"content"`
}
