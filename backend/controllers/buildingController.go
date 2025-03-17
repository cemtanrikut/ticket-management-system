package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"main.go/config"
	"main.go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetBuildings tüm binaları döndürür
func GetBuildings(w http.ResponseWriter, r *http.Request) {
	collection := config.GetCollection("buildings")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching buildings", http.StatusInternalServerError)
		return
	}

	var buildings []models.Building
	if err := cursor.All(context.TODO(), &buildings); err != nil {
		http.Error(w, "Error decoding buildings", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(buildings)
}

// CreateBuilding yeni bina oluşturur
func CreateBuilding(w http.ResponseWriter, r *http.Request) {
	var building models.Building
	if err := json.NewDecoder(r.Body).Decode(&building); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	building.ID = primitive.NewObjectID()
	building.CreatedAt = time.Now()

	collection := config.GetCollection("buildings")
	_, err := collection.InsertOne(context.TODO(), building)
	if err != nil {
		http.Error(w, "Error creating building", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(building)
}
