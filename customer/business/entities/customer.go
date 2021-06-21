package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID       primitive.ObjectID `bson:"_id"`
	Document string             `bson:"document"`
	Name     string             `bson:"name"`
	BirthDay time.Time          `bson:"birthDay"`
}

func NewCustomer(document, name string, birthDay time.Time) *Customer {
	return &Customer{
		ID:       primitive.NewObjectID(),
		Document: document,
		Name:     name,
		BirthDay: birthDay,
	}
}
