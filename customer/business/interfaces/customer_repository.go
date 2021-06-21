package interfaces

import "github.com/jailtonjunior94/bank/customer/business/entities"

type ICustomerRepository interface {
	Add(c *entities.Customer) (customer *entities.Customer, err error)
	GetByDocument(document string) (customer *entities.Customer, err error)
}
