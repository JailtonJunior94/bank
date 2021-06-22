package services

import (
	"encoding/json"

	"github.com/jailtonjunior94/bank/loan/business/dtos"
	"github.com/jailtonjunior94/bank/loan/business/entities"
	"github.com/jailtonjunior94/bank/loan/business/events"
	"github.com/jailtonjunior94/bank/loan/business/interfaces"
	"github.com/jailtonjunior94/bank/loan/infra/environments"
)

type LoanService struct {
	LoanRepository interfaces.ILoanRepository
	CustomerFacade interfaces.ICustomerFacade
	RabbitMQ       interfaces.IRabbitMQ
}

func NewLoanService(r interfaces.ILoanRepository, f interfaces.ICustomerFacade, b interfaces.IRabbitMQ) interfaces.ILoanService {
	return &LoanService{LoanRepository: r, CustomerFacade: f, RabbitMQ: b}
}

func (s *LoanService) LoanByDocument(document string) *dtos.HttpResponse {
	return nil
}

func (s *LoanService) CreateLoan(r *dtos.CreateLoanCommand) *dtos.HttpResponse {
	c, err := s.CustomerFacade.CustomerByDocument(r.Document)
	if err != nil {
		return dtos.BadRequest("Não foi possível cadastrar empréstimo")
	}

	newLoan := entities.NewLoan(c.Document, r.Income, r.Value, r.Rate, r.Quantity)
	l, err := s.LoanRepository.Add(newLoan)
	if err != nil {
		return dtos.ServerError()
	}

	loanMade := events.NewLoanMade(l.Document, l.Quantity, l.Rate, l.Value)
	body, err := json.Marshal(loanMade)
	if err != nil {
		return dtos.ServerError()
	}

	if err := s.RabbitMQ.SendMessage(environments.QueueLoanRisk, string(body)); err != nil {
		return dtos.ServerError()
	}

	res := dtos.NewLoanResponse(l.ID.Hex(), l.Document, l.Status, l.Income, l.Value, l.Rate, l.Quantity)
	return dtos.Created(res)
}
