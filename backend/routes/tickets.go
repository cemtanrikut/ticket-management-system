package routes

import (
	"ticket-management-system/controllers"
	"ticket-management-system/middleware"

	"github.com/gorilla/mux"
)

// RegisterTicketRoutes ticket işlemleri için route'ları ekler
func RegisterTicketRoutes(router *mux.Router) {
	ticketRouter := router.PathPrefix("/tickets").Subrouter()
	ticketRouter.Use(middleware.AuthMiddleware)

	ticketRouter.HandleFunc("", controllers.GetTickets).Methods("GET")
	ticketRouter.HandleFunc("/{id}", controllers.GetTicketByID).Methods("GET")
	ticketRouter.HandleFunc("", controllers.CreateTicket).Methods("POST")
	ticketRouter.HandleFunc("/{id}", controllers.UpdateTicket).Methods("PUT")
	ticketRouter.HandleFunc("/{id}", controllers.DeleteTicket).Methods("DELETE")
}
