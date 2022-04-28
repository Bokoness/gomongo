package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type customClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func CreateCookieToken(uid primitive.ObjectID, w http.ResponseWriter) {
	expireDays := 3
	d, _ := time.ParseDuration(fmt.Sprintf("%dh", expireDays*24))
	claims := customClaims{
		Uid: uid.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hash, e := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if e != nil {
		http.Error(w, "Cant log in", 500)
	}
	c := http.Cookie{Name: "uid", Value: hash, Path: "/"}
	http.SetCookie(w, &c)
}

func ClearAuth(w http.ResponseWriter) {
	c := http.Cookie{
		Name:   "uid",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
}

func BodyIntoMap(r *http.Request) (map[string]string, error) {
	m := make(map[string]string)
	e := json.NewDecoder(r.Body).Decode(&m)
	if e != nil {
		return nil, e
	}
	return m, nil
}

func WriteDataIntoResponse(d interface{}, w http.ResponseWriter) {
	j, e := json.Marshal(d)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(j)
}
