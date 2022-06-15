package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var secret_token string

func AuthUser(r *http.Request) bool {
	tok, err := r.Cookie("Authorization Token")
	if err != nil {
		return false
	}

	return verifyToken(tok.Value)
}

func verifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println(ok)
		}
		setSecretToken()
		return []byte(secret_token), nil
	})

	if err != nil {
		log.Println(err)
	}

	return token.Valid
}

func GenerateToken(w http.ResponseWriter) {
	new_token := jwt.New(jwt.SigningMethodHS256)

	setSecretToken()
	tokenString, err := new_token.SignedString([]byte(secret_token))
	if err != nil {
		log.Println(err)
	}
	http.SetCookie(w, &http.Cookie{Name: "Authorization Token", Value: tokenString, MaxAge: 86400, Secure: true, HttpOnly: true})
}

func setSecretToken() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}

	secret_token = os.Getenv("JWT_TOKEN")
}
