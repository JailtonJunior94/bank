package ioc

import (
	"github.com/jailtonjunior94/bank/risk/business/interfaces"
	"github.com/jailtonjunior94/bank/risk/business/services"
	"github.com/jailtonjunior94/bank/risk/infra/bus"
	"github.com/jailtonjunior94/bank/risk/infra/environments"
	"github.com/jailtonjunior94/bank/risk/infra/facades"
	"github.com/jailtonjunior94/bank/risk/presentation/handlers"
)

var (
	RabbitMQ         interfaces.IRabbitMQ
	EvaluationFacade interfaces.IEvaluationFacade
	LoanFacade       interfaces.ILoanFacade
	RiskService      interfaces.IRiskService
	RiskHandle       *handlers.RiskHandler
)

func SetupDependencyInjection() {
	RabbitMQ = bus.NewRabbitMQ(environments.RabbitMQConnection)
	EvaluationFacade = facades.NewEvaluationFacade()
	LoanFacade = facades.NewLoanFacade()
	RiskService = services.NewRiskService(LoanFacade, EvaluationFacade)
	RiskHandle = handlers.NewRiskHandler(RiskService)
}
