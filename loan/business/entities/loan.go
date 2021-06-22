package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Pending  = "PENDING"
	Approved = "APPROVED"
	Refused  = "REFUSED"
)

type Loan struct {
	ID       primitive.ObjectID `bson:"_id"`
	Document string             `bson:"document"`
	Income   float64            `bson:"income"`
	Value    float64            `bson:"value"`
	Quantity int                `bson:"quantity"`
	Rate     float64            `bson:"rate"`
	Status   string             `bson:"status"`
}

func NewLoan(document string, income, value, rate float64, quantity int) *Loan {
	return &Loan{
		ID:       primitive.NewObjectID(),
		Document: document,
		Income:   income,
		Value:    value,
		Quantity: quantity,
		Rate:     rate,
		Status:   Pending,
	}
}

func (l *Loan) ChangeStatus(status string) {
	l.Status = status
}
