package facades

import (
	"math/rand"
	"time"

	"github.com/jailtonjunior94/bank/risk/business/dtos"
	"github.com/jailtonjunior94/bank/risk/business/interfaces"
)

type Rating int

const (
	A Rating = iota + 1 // Super Potencial = 1
	B                   // Médio Potencial = 2
	C                   // Baixo Potencial = 3
	D                   // Negativado = 4
)

type EvaluationFacade struct {
}

func NewEvaluationFacade() interfaces.IEvaluationFacade {
	return &EvaluationFacade{}
}

func (f *EvaluationFacade) ToEvaluate(l *dtos.Loan) (e *dtos.Evaluation, err error) {
	rating := f.getRating()

	switch {
	case rating == D:
		return dtos.NewEvaluation(false, "Cliente negativado"), nil
	case rating == A:
		return dtos.NewEvaluation(true, "Cliente aprovado, devido a super potêncial"), nil
	case rating == B || l.Income > 1000 || l.Quantity < 5 || l.Value < 200:
		return dtos.NewEvaluation(true, "Cliente aprovado"), nil
	case rating == C && l.Income > 1200 || l.Quantity < 5 || l.Value < 400:
		return dtos.NewEvaluation(true, "Cliente aprovado"), nil
	default:
		return dtos.NewEvaluation(false, "Cliente com baixo potêncial"), nil
	}
}

func (f *EvaluationFacade) getRating() Rating {
	time.Sleep(10 * time.Microsecond)

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 4
	rating := rand.Intn(max-min+1) + min

	switch {
	case rating == 1:
		return A
	case rating == 2:
		return B
	case rating == 3:
		return C
	case rating == 4:
		return D
	}

	return 0
}
