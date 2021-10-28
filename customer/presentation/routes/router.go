package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	AddCustomerRouter(v1)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Healthy",
			"time":    time.Now(),
		})
	})
}
