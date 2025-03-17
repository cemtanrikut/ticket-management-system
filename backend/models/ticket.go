package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Ticket struct MongoDB için ticket modelini temsil eder
type Ticket struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"` // to-do, in-progress, done
	CustomerID  primitive.ObjectID `bson:"customerId,omitempty" json:"customerId,omitempty"`
	BuildingID  primitive.ObjectID `bson:"buildingId,omitempty" json:"buildingId,omitempty"`
	WorkerID    primitive.ObjectID `bson:"workerId,omitempty" json:"workerId,omitempty"`
	CreatedBy   primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
