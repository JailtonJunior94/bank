package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jailtonjunior94/bank/customer/infra/database"
	"github.com/jailtonjunior94/bank/customer/infra/environments"
	"github.com/jailtonjunior94/bank/customer/infra/ioc"
	"github.com/jailtonjunior94/bank/customer/presentation/routes"

	healthcheck "github.com/aschenmaker/fiber-health-check"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(healthcheck.New())

	environments.NewSettings()

	mongoConnection := database.NewConnection()
	defer mongoConnection.Disconnect()

	ioc.SetupDependencyInjection(mongoConnection)
	routes.SetupRoutes(app)

	fmt.Printf("🚀 Customer API is running on http://localhost:%v", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
