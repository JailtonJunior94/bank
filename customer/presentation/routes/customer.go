package routes

import (
	"github.com/jailtonjunior94/bank/customer/infra/ioc"

	"github.com/gofiber/fiber/v2"
)

func AddCustomerRouter(router fiber.Router) {
	router.Get("/customers/:document", ioc.CustomerController.CustomerByDocument)
	router.Post("/customers", ioc.CustomerController.CreateCustomer)
}
