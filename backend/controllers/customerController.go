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

// GetCustomers tüm müşterileri döndürür
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	collection := config.GetCollection("customers")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching customers", http.StatusInternalServerError)
		return
	}

	var customers []models.Customer
	if err := cursor.All(context.TODO(), &customers); err != nil {
		http.Error(w, "Error decoding customers", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customers)
}

// CreateCustomer yeni müşteri oluşturur
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	customer.ID = primitive.NewObjectID()
	customer.CreatedAt = time.Now()

	collection := config.GetCollection("customers")
	_, err := collection.InsertOne(context.TODO(), customer)
	if err != nil {
		http.Error(w, "Error creating customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

// GetCustomerByID belirli bir müşteriyi döndürür
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := config.GetCollection("customers")
	var customer models.Customer
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&customer)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomer müşteriyi günceller
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var updateData models.Customer
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	collection := config.GetCollection("customers")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updateData})
	if err != nil {
		http.Error(w, "Error updating customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer updated"})
}

// DeleteCustomer müşteriyi siler
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := config.GetCollection("customers")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, "Error deleting customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted"})
}
