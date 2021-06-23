package interfaces

import "github.com/jailtonjunior94/bank/loan/business/entities"

type ILoanRepository interface {
	Add(l *entities.Loan) (loan *entities.Loan, err error)
	GetById(id string) (loan *entities.Loan, err error)
	GetByDocument(document string) (loans []entities.Loan, err error)
	Update(l *entities.Loan) (loan *entities.Loan, err error)
}
