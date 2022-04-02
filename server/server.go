package server

import (
	"fmt"
	"gomongo/server/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func LunchServer() {
	r := mux.NewRouter()
	routes.TodoRoutes(r)
	routes.AuthRoutes(r)
	// r.Use(middleware.ServerHeaders)
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Server is running on port %s", port)
	http.ListenAndServe(port, r)
}
