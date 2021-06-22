package routes

import (
	"github.com/jailtonjunior94/bank/loan/infra/ioc"

	"github.com/gofiber/fiber/v2"
)

func AddLoanRouter(router fiber.Router) {
	router.Get("/loans/:document", ioc.LoanController.LoanByDocument)
	router.Post("/loans", ioc.LoanController.CreateLoan)
}
