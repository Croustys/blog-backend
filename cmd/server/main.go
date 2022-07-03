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

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func setSecretToken() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
}
