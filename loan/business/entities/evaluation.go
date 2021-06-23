package entities

type Evaluation struct {
	Approved    bool   `bson:"approved"`
	Description string `bson:"description"`
}

func NewEvaluation(approved bool, description string) *Evaluation {
	return &Evaluation{
		Approved:    approved,
		Description: description,
	}
}
