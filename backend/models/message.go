package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message struct MongoDB için mesaj modelini temsil eder
type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TicketID  primitive.ObjectID `bson:"ticketId" json:"ticketId"`
	SenderID  primitive.ObjectID `bson:"senderId" json:"senderId"`
	Text      string             `bson:"text" json:"text"`
	FileURL   string             `bson:"fileUrl,omitempty" json:"fileUrl,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}
