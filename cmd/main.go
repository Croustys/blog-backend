package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Croustys/blog-backend/internal/controller"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")
	setSecretToken()
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/create", controller.CreatePost)
	r.Path("/posts").Queries("offset", "{offset:[0-9]+}", "limit", "{limit:[0-9]+}").HandlerFunc(controller.GetPostsLazy)
	r.HandleFunc("/posts", controller.GetPosts)
	r.HandleFunc("/post/{id}", controller.GetPost)
	r.HandleFunc("/ping", controller.Ping)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func setSecretToken() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
}
