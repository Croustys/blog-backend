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

	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/create", controller.CreatePost).Methods("POST")
	r.Path("/posts").Queries("offset", "{offset:[0-9]+}", "limit", "{limit:[0-9]+}").HandlerFunc(controller.GetPostsLazy).Methods("GET")
	r.HandleFunc("/posts", controller.GetPosts).Methods("GET")
	r.HandleFunc("/post/{id}", controller.GetPost).Methods("GET")
	r.HandleFunc("/ping", controller.Ping).Methods("GET")
	r.HandleFunc("/user/{username}", controller.GetUserPosts).Methods("GET")
	r.HandleFunc("/user", controller.GetUser).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"https://quoteshare.vercel.app", "https://blog.barabasakos.hu"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headersOk, originsOk, methodsOk, handlers.AllowCredentials())(r)))
}

func setSecretToken() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
}
