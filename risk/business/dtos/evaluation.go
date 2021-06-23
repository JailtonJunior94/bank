package dtos

type Evaluation struct {
	Approved    bool   `json:"approved"`
	Description string `json:"description"`
}

func NewEvaluation(approved bool, description string) *Evaluation {
	return &Evaluation{
		Approved:    approved,
		Description: description,
	}
}
