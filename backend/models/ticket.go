package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Ticket struct MongoDB için ticket modelini temsil eder
type Ticket struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TicketID         string             `bson:"ticketId" json:"ticketId"`
	Title            string             `bson:"title" json:"title"`
	Description      string             `bson:"description" json:"description"`
	Status           string             `bson:"status" json:"status"` // to-do, in-progress, done
	WorkerID         primitive.ObjectID `bson:"workerId,omitempty" json:"workerId,omitempty"`
	BuildingID       primitive.ObjectID `bson:"buildingId,omitempty" json:"buildingId,omitempty"`
	CustomerID       primitive.ObjectID `bson:"customerId,omitempty" json:"customerId,omitempty"`
	NotificationType string             `bson:"notificationType" json:"notificationType"`
	CreatedFrom      string             `bson:"createdFrom" json:"createdFrom"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
}
