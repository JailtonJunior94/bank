package interfaces

import "github.com/jailtonjunior94/bank/loan/business/dtos"

type ILoanService interface {
	LoanByDocument(document string) *dtos.HttpResponse
	CreateLoan(request *dtos.CreateLoanCommand) *dtos.HttpResponse
}
