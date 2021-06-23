package handlers

import (
	"encoding/json"

	"github.com/jailtonjunior94/bank/risk/business/dtos"
	"github.com/jailtonjunior94/bank/risk/business/interfaces"
)

type RiskHandler struct {
	Service interfaces.IRiskService
}

func NewRiskHandler(s interfaces.IRiskService) *RiskHandler {
	return &RiskHandler{Service: s}
}

func (s *RiskHandler) AnalyzeRisk(data []byte) error {
	var loan dtos.Loan
	if err := json.Unmarshal(data, &loan); err != nil {
		return err
	}

	if err := s.Service.AnalyzeRisk(&loan); err != nil {
		return err
	}

	return nil
}
