package controllers

import (
	"github.com/jailtonjunior94/bank/customer/business/dtos"
	"github.com/jailtonjunior94/bank/customer/business/interfaces"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	Service interfaces.ICustomerService
}

func NewCustomerController(s interfaces.ICustomerService) *CustomerController {
	return &CustomerController{Service: s}
}

func (s *CustomerController) CustomerByDocument(c *fiber.Ctx) error {
	response := s.Service.CustomerByDocument(c.Params("document"))
	return c.Status(response.StatusCode).JSON(response.Data)
}

func (s *CustomerController) CreateCustomer(c *fiber.Ctx) error {
	request := new(dtos.CreateCustomerCommand)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := request.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := s.Service.CreateCustomer(request)
	return c.Status(response.StatusCode).JSON(response.Data)
}
