package main

import (
	"blog-backend/internal/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")
	setSecretToken()
	r := mux.NewRouter()

	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/create", controller.CreatePost)
	r.Path("/posts").Queries("offset", "{offset:[0-9]+}", "limit", "{limit:[0-9]+}").HandlerFunc(controller.GetPostsLazy)
	r.HandleFunc("/posts", controller.GetPosts)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func setSecretToken() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
}
