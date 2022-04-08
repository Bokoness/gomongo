package user

import (
	"encoding/json"
	"gomongo/db/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var u models.User
	u.FindById(params["id"])
	jsonData, err := json.Marshal(u.AsMapNoPwd())
	if err != nil {
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var u models.User
	u.FindByIdAndUpdate(params["id"])
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var u models.User
	err := u.FindById(params["id"])
	if err != nil {
		w.WriteHeader(401)
	}
	err = u.Destroy()
	if err != nil {
		w.WriteHeader(401)
	}
}
