package events

type LoanMade struct {
	LoanID   string  `json:"loanId"`
	Income   float64 `json:"income"`
	Rate     float64 `json:"rate"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
}

func NewLoanMade(loanID string, quantity int, income, rate, value float64) *LoanMade {
	return &LoanMade{
		LoanID:   loanID,
		Income:   income,
		Rate:     rate,
		Value:    value,
		Quantity: quantity,
	}
}
