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

// GetWorkers tüm çalışanları döndürür
func GetWorkers(w http.ResponseWriter, r *http.Request) {
	collection := config.GetCollection("workers")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching workers", http.StatusInternalServerError)
		return
	}

	var workers []models.Worker
	if err := cursor.All(context.TODO(), &workers); err != nil {
		http.Error(w, "Error decoding workers", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(workers)
}

// CreateWorker yeni çalışan oluşturur
func CreateWorker(w http.ResponseWriter, r *http.Request) {
	var worker models.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	worker.ID = primitive.NewObjectID()
	worker.CreatedAt = time.Now()

	collection := config.GetCollection("workers")
	_, err := collection.InsertOne(context.TODO(), worker)
	if err != nil {
		http.Error(w, "Error creating worker", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(worker)
}
