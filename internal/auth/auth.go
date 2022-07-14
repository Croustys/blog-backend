package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var secret_token string

func AuthUser(r *http.Request) bool {
	tok, err := r.Cookie("AuthToken")
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

func GenerateToken(w http.ResponseWriter, email string, username string) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &Claims{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	new_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	setSecretToken()
	tokenString, err := new_token.SignedString([]byte(secret_token))
	if err != nil {
		log.Println(err)
	}
	http.SetCookie(w, &http.Cookie{Name: "AuthToken", Value: tokenString, MaxAge: 86400, Secure: false, HttpOnly: true, Path: "/"})
}

func GetPayload(r *http.Request) (string, string) {
	tok, err := r.Cookie("AuthToken")
	if err != nil {
		log.Println(err)
	}

	claims := &Claims{}
	_, err = jwt.ParseWithClaims(tok.Value, claims, func(t *jwt.Token) (interface{}, error) {
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

	return claims.Email, claims.Username
}

func RemoveToken(w http.ResponseWriter, r *http.Request) {
	currentCookie, err := r.Cookie("AuthToken")
	if err != nil {
		log.Println(err)
	}
	currentCookie.MaxAge = -1
	http.SetCookie(w, currentCookie)
}

func setSecretToken() {
	secret_token = os.Getenv("JWT_TOKEN")
}
