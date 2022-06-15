package controller

import (
	"blog-backend/internal/auth"
	"net/http"
)

type UserS struct {
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	/* var u UserS
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} */

	if auth.AuthUser(r) {
		//statusok
	}
	auth.GenerateToken(w)
}
