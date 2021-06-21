package ioc

import (
	"github.com/jailtonjunior94/bank/customer/business/interfaces"
	"github.com/jailtonjunior94/bank/customer/business/services"
	"github.com/jailtonjunior94/bank/customer/infra/database"
	"github.com/jailtonjunior94/bank/customer/infra/repositories"
	"github.com/jailtonjunior94/bank/customer/presentation/controllers"
)

var (
	MongoConnection    database.IMongoConnection
	CustomerRepository interfaces.ICustomerRepository
	CustomerService    interfaces.ICustomerService
	CustomerController *controllers.CustomerController
)

func SetupDependencyInjection(mongoConnection database.IMongoConnection) {
	MongoConnection = mongoConnection
	CustomerRepository = repositories.NewCustomerRepository(MongoConnection)
	CustomerService = services.NewCustomerService(CustomerRepository)
	CustomerController = controllers.NewCustomerController(CustomerService)
}
