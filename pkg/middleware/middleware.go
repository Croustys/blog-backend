package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Croustys/blog-backend/internal/auth"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

var allowedSlugs []string = []string{"/login", "/register", "/ping", "/posts"}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := r.URL
		if contains(allowedSlugs, slug.Path) {
			next.ServeHTTP(w, r)
		} else if auth.AuthUser(r) {
			next.ServeHTTP(w, r)
		} else {
			json, err := json.Marshal(map[string]string{"StatusMessage": "Unauthorized"})
			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(json)
		}
	})
}
