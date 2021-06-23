package interfaces

import "github.com/jailtonjunior94/bank/risk/business/dtos"

type IEvaluationFacade interface {
	ToEvaluate(l *dtos.Loan) (e *dtos.Evaluation, err error)
}
