package user

import (
	"gomongo/db/models"
	"gomongo/services"
	"net/http"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	var u models.User
	u.FindById(services.ParseIdFromReq(r))
	u.Destroy()
}
