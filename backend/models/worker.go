package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Worker struct MongoDB için çalışan modelini temsil eder
type Worker struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	CustomerID primitive.ObjectID `bson:"customerId,omitempty" json:"customerId,omitempty"`
	Phone      string             `bson:"phone" json:"phone"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"-"`
	Department string             `bson:"department" json:"department"`
	Role       string             `bson:"role" json:"role"`
	StartDate  time.Time          `bson:"startDate" json:"startDate"`
	Status     string             `bson:"status" json:"status"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
}
