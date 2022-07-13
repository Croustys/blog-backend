package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Croustys/blog-backend/internal/auth"
	"github.com/Croustys/blog-backend/internal/db"
	"github.com/Croustys/blog-backend/internal/types"

	"github.com/gorilla/mux"
)

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

	var u types.UserS
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		unAuthHttpResponse(&w, "Missing Credentials")
		return
	}

	isLoginSuccessful, username := db.LoginUser(u.Email, u.Password)
	if !isLoginSuccessful {
		unAuthHttpResponse(&w, "Wrong Credentials")
		return
	}

	auth.GenerateToken(w, u.Email, username)

	json, err := json.Marshal(map[string]string{"StatusMessage": "Login Successful"})
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u types.UserS
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success := db.RegisterUser(u.Email, u.Password, u.Username)

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

	var p types.PostS
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
	}

	authorEmail, authorUsername := auth.GetPayload(r)
	if !db.SavePost(authorEmail, authorUsername, p.Title, p.Content) {
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
func GetPostsLazy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	offset := mux.Vars(r)["offset"]
	limit := mux.Vars(r)["limit"]
	offsetInt, _ := strconv.ParseInt(offset, 10, 64)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)

	data := db.GetPosts(offsetInt, limitInt)

	json, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := db.GetPosts(0, 0)
	json, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	data := db.GetPost(id)

	json, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := mux.Vars(r)["username"]
	data := db.GetPostsByUsername(username)
	json, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authorEmail, authorUsername := auth.GetPayload(r)

	var resp types.GetUser

	resp.Posts = db.GetPostsByUsername(authorUsername)
	resp.Username = authorUsername
	resp.Email = authorEmail

	json, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
