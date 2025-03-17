package routes

import (
	"ticket-management-system/controllers"

	"github.com/gorilla/mux"
)

// RegisterAuthRoutes authentication için gerekli route'ları ekler
func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
}
