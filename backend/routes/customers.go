package routes

import (
	"main.go/middleware"

	"main.go/controllers"

	"github.com/gorilla/mux"
)

// RegisterCustomerRoutes müşteri işlemleri için route'ları ekler
func RegisterCustomerRoutes(router *mux.Router) {
	customerRouter := router.PathPrefix("/customers").Subrouter()
	customerRouter.Use(middleware.AuthMiddleware)

	customerRouter.HandleFunc("", controllers.GetCustomers).Methods("GET")
	customerRouter.HandleFunc("/{id}", controllers.GetCustomerByID).Methods("GET")
	customerRouter.HandleFunc("", controllers.CreateCustomer).Methods("POST")
	customerRouter.HandleFunc("/{id}", controllers.UpdateCustomer).Methods("PUT")
	customerRouter.HandleFunc("/{id}", controllers.DeleteCustomer).Methods("DELETE")
}
