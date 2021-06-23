package interfaces

import "github.com/jailtonjunior94/bank/risk/business/dtos"

type ILoanFacade interface {
	UpdateLoan(id string, r *dtos.Evaluation) error
}
