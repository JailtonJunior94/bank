package dtos

import (
	"github.com/jailtonjunior94/bank/loan/business/entities"
)

type LoanResponse struct {
	ID         string              `json:"id"`
	Income     float64             `json:"income"`
	Value      float64             `json:"value"`
	Quantity   int                 `json:"quantity"`
	Rate       float64             `json:"rate"`
	Status     string              `json:"status"`
	Customer   *CustomerResponse   `json:"customer,omitempty"`
	Evaluation *EvaluationResponse `json:"evaluation,omitempty"`
}

func NewLoanResponse(id, status string, income, value, rate float64, quantity int) *LoanResponse {
	return &LoanResponse{
		ID:       id,
		Income:   income,
		Value:    value,
		Quantity: quantity,
		Rate:     rate,
		Status:   status,
	}
}

func (r *LoanResponse) AddCustomer(c *entities.Customer) {
	if c != nil {
		r.Customer = &CustomerResponse{
			Document: c.Document,
			Name:     c.Name,
		}
	}
}

func (r *LoanResponse) AddEvaluation(e *entities.Evaluation) {
	if e != nil {
		r.Evaluation = &EvaluationResponse{
			Approved:    e.Approved,
			Description: e.Description,
		}
	}
}
