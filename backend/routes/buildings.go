package routes

import (
	"ticket-management-system/controllers"
	"ticket-management-system/middleware"

	"github.com/gorilla/mux"
)

// RegisterBuildingRoutes bina işlemleri için route'ları ekler
func RegisterBuildingRoutes(router *mux.Router) {
	buildingRouter := router.PathPrefix("/buildings").Subrouter()
	buildingRouter.Use(middleware.AuthMiddleware)

	buildingRouter.HandleFunc("", controllers.GetBuildings).Methods("GET")
	buildingRouter.HandleFunc("/{id}", controllers.GetBuildingByID).Methods("GET")
	buildingRouter.HandleFunc("", controllers.CreateBuilding).Methods("POST")
	buildingRouter.HandleFunc("/{id}", controllers.UpdateBuilding).Methods("PUT")
	buildingRouter.HandleFunc("/{id}", controllers.DeleteBuilding).Methods("DELETE")
}
