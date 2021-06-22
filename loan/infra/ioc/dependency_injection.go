package ioc

import (
	"github.com/jailtonjunior94/bank/loan/business/interfaces"
	"github.com/jailtonjunior94/bank/loan/business/services"
	"github.com/jailtonjunior94/bank/loan/infra/bus"
	"github.com/jailtonjunior94/bank/loan/infra/database"
	"github.com/jailtonjunior94/bank/loan/infra/environments"
	"github.com/jailtonjunior94/bank/loan/infra/facades"
	"github.com/jailtonjunior94/bank/loan/infra/repositories"
	"github.com/jailtonjunior94/bank/loan/presentation/controllers"
)

var (
	MongoConnection database.IMongoConnection
	LoanRepository  interfaces.ILoanRepository
	CustomerFacade  interfaces.ICustomerFacade
	RabbitMQ        interfaces.IRabbitMQ
	LoanService     interfaces.ILoanService
	LoanController  *controllers.LoanController
)

func SetupDependencyInjection(mongoConnection database.IMongoConnection) {
	MongoConnection = mongoConnection
	LoanRepository = repositories.NewLoanRepository(MongoConnection)
	CustomerFacade = facades.NewCustomerFacade()
	RabbitMQ = bus.New(environments.RabbitMQConnection)
	LoanService = services.NewLoanService(LoanRepository, CustomerFacade, RabbitMQ)
	LoanController = controllers.NewLoanController(LoanService)
}
