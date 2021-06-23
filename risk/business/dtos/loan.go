package dtos

type Loan struct {
	Document string  `json:"document"`
	Income   float64 `json:"income"`
	Rate     float64 `json:"rate"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
}
