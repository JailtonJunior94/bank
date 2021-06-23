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
	loans, err := s.LoanRepository.GetByDocument(document)
	if err != nil {
		return dtos.ServerError()
	}

	if len(loans) == 0 {
		return dtos.NotFound("Não foi encontrado nenhum empréstimo")
	}

	var res []*dtos.LoanResponse
	for _, l := range loans {
		loan := dtos.NewLoanResponse(l.ID.Hex(), l.Status, l.Income, l.Value, l.Rate, l.Quantity)
		loan.AddCustomer(l.Customer)
		loan.AddEvaluation(l.Evaluation)
		res = append(res, loan)
	}

	return dtos.Ok(res)
}

func (s *LoanService) CreateLoan(r *dtos.CreateLoanCommand) *dtos.HttpResponse {
	c, err := s.CustomerFacade.CustomerByDocument(r.Document)
	if err != nil {
		return dtos.BadRequest("Não foi possível cadastrar empréstimo")
	}

	newLoan := entities.NewLoan(r.Income, r.Value, r.Rate, r.Quantity)
	newLoan.AddCustomer(entities.NewCustomer(c.Document, c.Name))
	l, err := s.LoanRepository.Add(newLoan)
	if err != nil {
		return dtos.ServerError()
	}

	loanMade := events.NewLoanMade(l.ID.Hex(), l.Quantity, l.Income, l.Rate, l.Value)
	body, err := json.Marshal(loanMade)
	if err != nil {
		return dtos.ServerError()
	}

	if err := s.RabbitMQ.SendMessage(environments.QueueLoanRisk, string(body)); err != nil {
		return dtos.ServerError()
	}

	res := dtos.NewLoanResponse(l.ID.Hex(), l.Status, l.Income, l.Value, l.Rate, l.Quantity)
	res.AddCustomer(l.Customer)
	res.AddEvaluation(l.Evaluation)

	return dtos.Created(res)
}

func (s *LoanService) UpdateLoan(id string, r *dtos.UpdateLoanCommand) *dtos.HttpResponse {
	l, err := s.LoanRepository.GetById(id)
	if err != nil {
		return dtos.ServerError()
	}

	if l == nil {
		return dtos.NotFound("Não foi encontrado nenhum empréstimo")
	}

	l.AddEvaluation(entities.NewEvaluation(r.Approved, r.Description))
	l.ChangeStatus(r.Approved)

	_, err = s.LoanRepository.Update(l)
	if err != nil {
		return dtos.ServerError()
	}

	return dtos.NoContent()
}
