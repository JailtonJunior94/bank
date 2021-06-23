package events

type LoanMade struct {
	Document string  `json:"document"`
	Income   float64 `json:"income"`
	Rate     float64 `json:"rate"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
}

func NewLoanMade(document string, quantity int, income, rate, value float64) *LoanMade {
	return &LoanMade{
		Document: document,
		Income:   income,
		Rate:     rate,
		Value:    value,
		Quantity: quantity,
	}
}
