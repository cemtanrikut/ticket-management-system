package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer struct MongoDB için müşteri modelini temsil eder
type Customer struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name                 string             `bson:"name" json:"name"`
	Address              string             `bson:"address" json:"address"`
	PostCode             string             `bson:"postCode" json:"postCode"`
	Plaats               string             `bson:"plaats" json:"plaats"`
	Country              string             `bson:"country" json:"country"`
	Phone                string             `bson:"phone" json:"phone"`
	Email                string             `bson:"email" json:"email"`
	Password             string             `bson:"password" json:"-"`
	Website              string             `bson:"website" json:"website"`
	Logo                 string             `bson:"logo" json:"logo"`
	Status               string             `bson:"status" json:"status"`
	Supplier             string             `bson:"supplier" json:"supplier"`
	BtwNumber            string             `bson:"btwNumber" json:"btwNumber"`
	Kvk                  string             `bson:"kvk" json:"kvk"`
	Vestigingsnummer     string             `bson:"vestigingsnummer" json:"vestigingsnummer"`
	Relatiebeheerder     string             `bson:"relatiebeheerder" json:"relatiebeheerder"`
	GlobalLocationNumber string             `bson:"globalLocationNumber" json:"globalLocationNumber"`
	Moederonderneming    string             `bson:"moederonderneming" json:"moederonderneming"`
	Remarks              string             `bson:"remarks" json:"remarks"`
	Contacts             []string           `bson:"contacts" json:"contacts"`
	CreatedAt            time.Time          `bson:"createdAt" json:"createdAt"`
}
