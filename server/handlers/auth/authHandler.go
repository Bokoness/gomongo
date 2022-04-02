package auth

import (
	"encoding/json"
	"fmt"
	"gomongo/db/models"
	"gomongo/services"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("u")
	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	u.Store()
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.User
	_ = json.NewDecoder(r.Body).Decode(&body)
	var u models.User
	u.FindByUsername(body.Username)
	if !services.ComparePasswords(u.Password, body.Password) {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
	}
	services.CreateCookieToken(u.Id, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	services.ClearAuth(w)
}
