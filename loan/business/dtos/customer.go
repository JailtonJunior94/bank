package dtos

import "time"

type CustomerResponse struct {
	ID       string    `json:"id,omitempty"`
	Document string    `json:"document,omitempty"`
	Name     string    `json:"name,omitempty"`
	BirthDay time.Time `json:"birthDay,omitempty"`
}

func NewCustomerResponse(id, document, name string, birthDay time.Time) *CustomerResponse {
	return &CustomerResponse{
		ID:       id,
		Document: document,
		Name:     name,
		BirthDay: birthDay,
	}
}
