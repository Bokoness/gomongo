package routes

import (
	"gomongo/server/handlers/user"
	"gomongo/server/middleware"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	s := r.PathPrefix("/todo").Subrouter()
	s.HandleFunc("/{id}", user.Destroy).Methods("DELETE")
	s.Use(middleware.UserAuth)
}
