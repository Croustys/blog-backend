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

func unAuthHttpResponse(w *http.ResponseWriter, msg string) {
	json, err := json.Marshal(map[string]string{"StatusMessage": msg})
	if err != nil {
		log.Println(err)
	}
	(*w).WriteHeader(http.StatusUnauthorized)
	(*w).Write(json)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if auth.AuthUser(r) { //return early if request has the correct jwt cookie
		json, err := json.Marshal(map[string]string{"StatusMessage": "Login Successful"})
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	var u UserS
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		unAuthHttpResponse(&w, "Missing Credentials")
		return
	}

	isLoginSuccessful := db.LoginUser(u.Email, u.Password)
	if !isLoginSuccessful {
		unAuthHttpResponse(&w, "Wrong Credentials")
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
func CreatePost(w http.ResponseWriter, r *http.Request) {
	//create post
}
