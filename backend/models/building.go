package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Building struct MongoDB için bina modelini temsil eder
type Building struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name"`
	CustomerID    primitive.ObjectID `bson:"customerId,omitempty" json:"customerId,omitempty"`
	Address       string             `bson:"address" json:"address"`
	HouseNo       string             `bson:"houseNo" json:"houseNo"`
	PostCode      string             `bson:"postCode" json:"postCode"`
	Plaats        string             `bson:"plaats" json:"plaats"`
	Status        string             `bson:"status" json:"status"`
	Note          string             `bson:"note" json:"note"`
	CalculateType string             `bson:"calculateType" json:"calculateType"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
}
