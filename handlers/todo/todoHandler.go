package todo

import (
	"encoding/json"
	"gomongo/db/models"
	"gomongo/middleware"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	_ = json.NewDecoder(r.Body).Decode(&t)
	u, _ := middleware.FetchUserFromCookie(r)
	t.User = u.ID
	t.Create()
}

func Index(w http.ResponseWriter, r *http.Request) {
	// u := middleware.FetchUserFromCtx(r)
	// j, e := json.Marshal(u.Todos)
	// if e != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
	// w.Write(j)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var todo models.Todo
	todo.FindById(id)
	j, e := json.Marshal(todo)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(j)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var t models.Todo
	if e := json.NewDecoder(r.Body).Decode(&t); e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	u, _ := middleware.FetchUserFromCookie(r)
	t.ID = vars["id"]
	t.User = u.ID
	t.Save()
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	u, _ := middleware.FetchUserFromCookie(r)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var t models.Todo
	t.FindById(id)
	if t.ID != u.ID {
		w.WriteHeader(http.StatusUnauthorized)
	}
	t.Destroy()
}