package auth

import (
	"encoding/json"
	"gomongo/db/models"
	"gomongo/services"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	err := u.Store()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	services.WriteDataIntoResponse(u.AsMap(true, false), w)
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
	services.CreateCookieToken(u.ID, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	services.ClearAuth(w)
}
