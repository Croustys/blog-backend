package controller

import (
	"blog-backend/internal/auth"
	"blog-backend/internal/db"
	"encoding/json"
	"log"
	"net/http"
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

	isLoginSuccessful := db.LoginUser(u.Email, u.Password)
	if !isLoginSuccessful {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if auth.AuthUser(r) {
		json, err := json.Marshal(map[string]string{"StatusMessage": "Login Successful"})
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	auth.GenerateToken(w)

	json, err := json.Marshal(map[string]string{"StatusMessage": "Login Successful"})
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u UserS
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success := db.RegisterUser(u.Email, u.Password)

	if !success {
		w.WriteHeader(http.StatusBadRequest)
		json, err := json.Marshal(map[string]string{"StatusMessage": "Register Unsuccessful"})

		if err != nil {
			log.Println(err)
		}
		w.Write(json)
		return
	}

	json, err := json.Marshal(map[string]string{"StatusMessage": "Register Successful"})
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
