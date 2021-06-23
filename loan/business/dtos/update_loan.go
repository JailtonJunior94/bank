package dtos

type UpdateLoanCommand struct {
	Approved    bool   `json:"approved"`
	Description string `json:"description"`
}
