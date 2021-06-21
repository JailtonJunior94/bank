package interfaces

import "github.com/jailtonjunior94/bank/customer/business/dtos"

type ICustomerService interface {
	CustomerByDocument(document string) *dtos.HttpResponse
	CreateCustomer(request *dtos.CreateCustomerCommand) *dtos.HttpResponse
}
