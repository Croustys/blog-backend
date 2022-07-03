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
type PostS struct {
	Title   string
	Content string
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

	auth.GenerateToken(w, u.Email)

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
	w.Header().Set("Content-Type", "application/json")

	var p PostS
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
	}

	authorEmail := auth.GetPayload(r)
	if !db.SavePost(authorEmail, p.Title, p.Content) {
		w.WriteHeader(http.StatusInternalServerError)
		json, err := json.Marshal(map[string]string{"StatusMessage": "Unsuccessful creation of a new Blog"})
		if err != nil {
			log.Println(err)
		}
		w.Write(json)
		return
	}

	json, err := json.Marshal(map[string]string{"StatusMessage": "Successful creation of a new Blog"})
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
