package interfaces

import "github.com/jailtonjunior94/bank/loan/business/dtos"

type ICustomerFacade interface {
	CustomerByDocument(document string) (customer *dtos.CustomerResponse, err error)
}
