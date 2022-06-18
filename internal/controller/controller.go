package controller

import (
	"blog-backend/internal/auth"
	"blog-backend/internal/db"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserS struct {
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u UserS
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	failedLogin := db.LoginUser(u.Email, u.Password)
	if failedLogin {
		log.Println("Bad credentials")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if auth.AuthUser(r) {
		json, err := json.Marshal("Login successful")
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}

	auth.GenerateToken(w)

	json, err := json.Marshal("Login successful")
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u UserS
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashed)

	auth.GenerateToken(w)
	success, result := db.Insert(u.Email, u.Password)

	w.Header().Set("Content-Type", "application/json")

	if !success {
		w.WriteHeader(http.StatusBadRequest)
		json, err := json.Marshal("Register Unsuccessful")
		if err != nil {
			log.Println(err)
		}
		w.Write(json)
	}

	json, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
