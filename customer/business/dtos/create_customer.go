package dtos

import (
	"errors"
	"time"
)

type CreateCustomerCommand struct {
	Document string    `json:"document"`
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birthDay"`
}

func (c *CreateCustomerCommand) IsValid() error {
	if c.BirthDay.IsZero() {
		return errors.New("A Data é obrigatória")
	}

	if c.Document == "" {
		return errors.New("O Documento é obrigatório")
	}

	if c.Name == "" {
		return errors.New("O Nome é obrigatório")
	}

	return nil
}
