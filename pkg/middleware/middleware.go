package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware")
		//@TODO: handle & auth routes
		next.ServeHTTP(w, r)
	})
}
