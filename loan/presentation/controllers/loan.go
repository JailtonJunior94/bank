package controllers

import (
	"github.com/jailtonjunior94/bank/loan/business/dtos"
	"github.com/jailtonjunior94/bank/loan/business/interfaces"

	"github.com/gofiber/fiber/v2"
)

type LoanController struct {
	Service interfaces.ILoanService
}

func NewLoanController(s interfaces.ILoanService) *LoanController {
	return &LoanController{Service: s}
}

func (s *LoanController) LoanByDocument(c *fiber.Ctx) error {
	response := s.Service.LoanByDocument(c.Params("document"))
	return c.Status(response.StatusCode).JSON(response.Data)
}

func (s *LoanController) CreateLoan(c *fiber.Ctx) error {
	request := new(dtos.CreateLoanCommand)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := request.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := s.Service.CreateLoan(request)
	return c.Status(response.StatusCode).JSON(response.Data)
}

func (s *LoanController) UpdateLoan(c *fiber.Ctx) error {
	request := new(dtos.UpdateLoanCommand)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	response := s.Service.UpdateLoan(c.Params("id"), request)
	return c.Status(response.StatusCode).JSON(response.Data)
}
