package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Croustys/blog-backend/internal/controller"
	"github.com/Croustys/blog-backend/pkg/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")
	setSecretToken()
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/create", controller.CreatePost)
	r.Path("/posts").Queries("offset", "{offset:[0-9]+}", "limit", "{limit:[0-9]+}").HandlerFunc(controller.GetPostsLazy)
	r.HandleFunc("/posts", controller.GetPosts)
	r.HandleFunc("/post/{id}", controller.GetPost)
	r.HandleFunc("/ping", controller.Ping)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"}) //@TODO: change to frontends host url
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}

func setSecretToken() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
}
