package main

import (
	"blog-backend/internal/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Login)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
