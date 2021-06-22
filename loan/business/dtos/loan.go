package dtos

type LoanResponse struct {
	ID       string  `json:"id"`
	Document string  `json:"document"`
	Income   float64 `json:"income"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
	Rate     float64 `json:"rate"`
	Status   string  `json:"status"`
}

func NewLoanResponse(id, document, status string, income, value, rate float64, quantity int) *LoanResponse {
	return &LoanResponse{
		ID:       id,
		Document: document,
		Income:   income,
		Value:    value,
		Quantity: quantity,
		Rate:     rate,
		Status:   status,
	}
}
