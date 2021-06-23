package services

import (
	"github.com/jailtonjunior94/bank/risk/business/dtos"
	"github.com/jailtonjunior94/bank/risk/business/interfaces"
)

type RiskService struct {
	EvaluationFacade interfaces.IEvaluationFacade
	LoanFacade       interfaces.ILoanFacade
}

func NewRiskService(l interfaces.ILoanFacade, e interfaces.IEvaluationFacade) interfaces.IRiskService {
	return &RiskService{LoanFacade: l, EvaluationFacade: e}
}

func (s *RiskService) AnalyzeRisk(l *dtos.Loan) error {
	e, err := s.EvaluationFacade.ToEvaluate(l)
	if err != nil {
		return err
	}

	if err := s.LoanFacade.UpdateLoan(l.LoanID, e); err != nil {
		return err
	}

	return nil
}
