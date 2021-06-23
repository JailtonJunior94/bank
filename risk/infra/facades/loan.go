package facades

import "github.com/jailtonjunior94/bank/risk/business/interfaces"

type LoanFacade struct {
}

func NewLoanFacade() interfaces.ILoanFacade {
	return &LoanFacade{}
}

func (f *LoanFacade) GetById(id string) (loan interface{}, err error) {
	return nil, nil
}
