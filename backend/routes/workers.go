package routes

import (
	"main.go/controllers"
	"main.go/middleware"

	"github.com/gorilla/mux"
)

// RegisterWorkerRoutes çalışan işlemleri için route'ları ekler
func RegisterWorkerRoutes(router *mux.Router) {
	workerRouter := router.PathPrefix("/workers").Subrouter()
	workerRouter.Use(middleware.AuthMiddleware)

	workerRouter.HandleFunc("", controllers.GetWorkers).Methods("GET")
	workerRouter.HandleFunc("/{id}", controllers.GetWorkerByID).Methods("GET")
	workerRouter.HandleFunc("", controllers.CreateWorker).Methods("POST")
	workerRouter.HandleFunc("/{id}", controllers.UpdateWorker).Methods("PUT")
	workerRouter.HandleFunc("/{id}", controllers.DeleteWorker).Methods("DELETE")
}
