package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"ticket-management-system/config"
	"ticket-management-system/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetTickets tüm ticket'ları döndürür
func GetTickets(w http.ResponseWriter, r *http.Request) {
	collection := config.GetCollection("tickets")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching tickets", http.StatusInternalServerError)
		return
	}

	var tickets []models.Ticket
	if err := cursor.All(context.TODO(), &tickets); err != nil {
		http.Error(w, "Error decoding tickets", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tickets)
}

// CreateTicket yeni ticket oluşturur
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ticket.ID = primitive.NewObjectID()
	ticket.CreatedAt = time.Now()

	collection := config.GetCollection("tickets")
	_, err := collection.InsertOne(context.TODO(), ticket)
	if err != nil {
		http.Error(w, "Error creating ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

// GetTicketByID belirli bir ticket'ı döndürür
func GetTicketByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := config.GetCollection("tickets")
	var ticket models.Ticket
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&ticket)
	if err != nil {
		http.Error(w, "Ticket not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ticket)
}

// UpdateTicket ticket'ı günceller
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var updateData models.Ticket
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	collection := config.GetCollection("tickets")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updateData})
	if err != nil {
		http.Error(w, "Error updating ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ticket updated"})
}

// DeleteTicket ticket'ı siler
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := config.GetCollection("tickets")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, "Error deleting ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ticket deleted"})
}
