package dtos

type EvaluationResponse struct {
	Approved    bool   `bson:"approved"`
	Description string `bson:"description"`
}

func NewEvaluationResponse(approved bool, description string) *EvaluationResponse {
	return &EvaluationResponse{
		Approved:    approved,
		Description: description,
	}
}
