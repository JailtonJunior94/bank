package events

type LoanMade struct {
	Document string  `json:"document"`
	Rate     float64 `json:"rate"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
}

func NewLoanMade(document string, quantity int, rate, value float64) *LoanMade {
	return &LoanMade{
		Document: document,
		Rate:     rate,
		Value:    value,
		Quantity: quantity,
	}
}
