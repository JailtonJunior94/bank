package dtos

import (
	"errors"
)

type CreateLoanCommand struct {
	Document string  `json:"document"`
	Income   float64 `json:"income"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
	Rate     float64 `json:"rate"`
}

func (c *CreateLoanCommand) IsValid() error {
	if c.Document == "" {
		return errors.New("O Documento é obrigatório")
	}

	if c.Income == 0 {
		return errors.New("A Renda é obrigatória")
	}

	if c.Value == 0 {
		return errors.New("O Valor do Empréstimo é obrigatório")
	}

	if c.Quantity == 0 {
		return errors.New("A quantidade de parcelas é obrigatório")
	}

	if c.Rate == 0 {
		return errors.New("A Taxa é obrigatória")
	}

	return nil
}
