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
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.TodoRoutes(r)
	// r.Use(middleware.ServerHeaders)
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Server is running on port %s", port)
	http.ListenAndServe(port, r)
}
