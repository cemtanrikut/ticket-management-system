package routes

import (
	"ticket-management-system/controllers"
	"ticket-management-system/middleware"

	"github.com/gorilla/mux"
)

// RegisterMessageRoutes mesaj işlemleri için route'ları ekler
func RegisterMessageRoutes(router *mux.Router) {
	messageRouter := router.PathPrefix("/messages").Subrouter()
	messageRouter.Use(middleware.AuthMiddleware)

	messageRouter.HandleFunc("/{ticketId}", controllers.GetMessages).Methods("GET")
	messageRouter.HandleFunc("/{ticketId}", controllers.SendMessage).Methods("POST")
}
