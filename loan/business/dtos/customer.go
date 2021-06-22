package dtos

import "time"

type CustomerResponse struct {
	ID       string    `json:"id"`
	Document string    `json:"document"`
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birthDay"`
}

func NewCustomerResponse(id, document, name string, birthDay time.Time) *CustomerResponse {
	return &CustomerResponse{
		ID:       id,
		Document: document,
		Name:     name,
		BirthDay: birthDay,
	}
}
