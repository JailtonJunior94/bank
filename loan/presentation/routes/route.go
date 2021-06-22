package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	AddLoanRouter(v1)
}
