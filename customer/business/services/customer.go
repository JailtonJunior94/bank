package services

import (
	"github.com/jailtonjunior94/bank/customer/business/dtos"
	"github.com/jailtonjunior94/bank/customer/business/entities"
	"github.com/jailtonjunior94/bank/customer/business/interfaces"
)

type CustomerService struct {
	CustomerRepository interfaces.ICustomerRepository
}

func NewCustomerService(r interfaces.ICustomerRepository) interfaces.ICustomerService {
	return &CustomerService{CustomerRepository: r}
}

func (s *CustomerService) CustomerByDocument(document string) *dtos.HttpResponse {
	c, err := s.CustomerRepository.GetByDocument(document)
	if err != nil {
		return dtos.ServerError()
	}

	if c == nil {
		return dtos.NotFound("Não foi encontrado nenhum cliente")
	}

	response := dtos.NewCustomerResponse(c.ID.Hex(), c.Document, c.Name, c.BirthDay)
	return dtos.Ok(response)
}

func (s *CustomerService) CreateCustomer(request *dtos.CreateCustomerCommand) *dtos.HttpResponse {
	newCustomer := entities.NewCustomer(request.Document, request.Name, request.BirthDay)
	c, err := s.CustomerRepository.Add(newCustomer)
	if err != nil {
		return dtos.ServerError()
	}

	response := dtos.NewCustomerResponse(c.ID.Hex(), c.Document, c.Name, c.BirthDay)
	return dtos.Created(response)
}
