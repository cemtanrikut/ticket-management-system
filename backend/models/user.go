package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct MongoDB için kullanıcı modelini temsil eder
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"-"`
	Role       string             `bson:"role" json:"role"` // admin, customer, worker
	CustomerID primitive.ObjectID `bson:"customerId,omitempty" json:"customerId,omitempty"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
}
