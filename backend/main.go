package main

import (
	"log"
	"net/http"
	"os"

	"ticket-management-system/config"
	"ticket-management-system/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Ortam değişkenlerini yükle
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// MongoDB bağlantısını başlat
	config.ConnectDB()

	// Yeni router oluştur
	r := mux.NewRouter()

	// API Route'larını ekle
	routes.RegisterAuthRoutes(r)
	routes.RegisterTicketRoutes(r)
	routes.RegisterMessageRoutes(r)
	routes.RegisterCustomerRoutes(r)
	routes.RegisterBuildingRoutes(r)
	routes.RegisterWorkerRoutes(r)

	// Sunucuyu başlat
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
