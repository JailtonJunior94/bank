package routes

import "github.com/gofiber/fiber/v2"

func AddCustomerRouter(router fiber.Router) {
	router.Get("/customers/:document", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Customer")
	})
	router.Post("/customers", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON("Customer created")
	})
}
