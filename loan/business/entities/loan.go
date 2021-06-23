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
	ID         primitive.ObjectID `bson:"_id"`
	Income     float64            `bson:"income"`
	Value      float64            `bson:"value"`
	Quantity   int                `bson:"quantity"`
	Rate       float64            `bson:"rate"`
	Status     string             `bson:"status"`
	Customer   *Customer          `bson:"customer"`
	Evaluation *Evaluation        `bson:"evaluation"`
}

func NewLoan(income, value, rate float64, quantity int) *Loan {
	return &Loan{
		ID:       primitive.NewObjectID(),
		Income:   income,
		Value:    value,
		Quantity: quantity,
		Rate:     rate,
		Status:   Pending,
	}
}

func (l *Loan) ChangeStatus(status bool) {
	if status {
		l.Status = Approved
		return
	}
	l.Status = Refused
}

func (l *Loan) AddCustomer(c *Customer) {
	l.Customer = c
}

func (l *Loan) AddEvaluation(e *Evaluation) {
	l.Evaluation = e
}
