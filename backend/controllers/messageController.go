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

// GetMessages belirli bir ticket için tüm mesajları döndürür
func GetMessages(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, _ := primitive.ObjectIDFromHex(params["ticketId"])

	collection := config.GetCollection("messages")
	cursor, err := collection.Find(context.TODO(), bson.M{"ticketId": ticketID})
	if err != nil {
		http.Error(w, "Error fetching messages", http.StatusInternalServerError)
		return
	}

	var messages []models.Message
	if err := cursor.All(context.TODO(), &messages); err != nil {
		http.Error(w, "Error decoding messages", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}

// SendMessage belirli bir ticket için yeni bir mesaj ekler
func SendMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, _ := primitive.ObjectIDFromHex(params["ticketId"])

	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	message.ID = primitive.NewObjectID()
	message.TicketID = ticketID
	message.CreatedAt = time.Now()

	collection := config.GetCollection("messages")
	_, err := collection.InsertOne(context.TODO(), message)
	if err != nil {
		http.Error(w, "Error saving message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}
