package session

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/miruts/iJobs/security/rndtoken"
	"net/http"
	"time"
)

// Create creates and sets session cookie
func Create(claims jwt.Claims, sessionID string, expires int, signingKey []byte, w http.ResponseWriter) {

	signedString, err := rndtoken.Generate(signingKey, claims)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	c := http.Cookie{
		Name:     sessionID,
		Expires:  time.Now().Add(time.Hour * time.Duration(expires)),
		Value:    signedString,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

// Valid validates client cookie value
func Valid(cookieValue string, signingKey []byte) (bool, error) {
	valid, err := rndtoken.Valid(cookieValue, signingKey)
	if err != nil || !valid {
		return false, errors.New("Invalid Session Cookie")
	}
	return true, nil
}

// Remove expires existing session
func Remove(sessionID string, w http.ResponseWriter) {
	c := http.Cookie{
		Name:    sessionID,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, &c)
}
func RemoveMock(sessionID string, w http.ResponseWriter) {
	c := http.Cookie{
		Name:    sessionID,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, &c)
}
