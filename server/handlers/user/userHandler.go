package user

import (
	"gomongo/db/models"
	"gomongo/services"
	"net/http"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var u models.User
	u.FindByIdAndUpdate(params["id"])
	// filter := map
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	var u models.User
	u.FindById(services.ParseIdFromReq(r))
	u.Destroy()
}
