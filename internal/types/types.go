package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserS struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type PostS struct {
	AuthorEmail string `json:"authorEmail"`
	Username    string `json:"username"`
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:"required"`
}
type PostResponse struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `json:"username"`
	Title    string             `json:"title"`
	Content  string             `json:"content"`
}
