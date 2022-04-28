package routes

import (
	"gomongo/handlers/auth"
	"gomongo/middleware"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", auth.Login).Methods("POST")
	s.HandleFunc("/register", auth.Register).Methods("POST")
	s.HandleFunc("/logout", auth.Logout).Methods("POST")
	s.Use(middleware.ServerHeaders)
}
