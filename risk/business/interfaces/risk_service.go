package interfaces

import "github.com/jailtonjunior94/bank/risk/business/dtos"

type IRiskService interface {
	AnalyzeRisk(l *dtos.Loan) error
}
