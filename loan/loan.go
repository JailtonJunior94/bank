package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jailtonjunior94/bank/loan/infra/database"
	"github.com/jailtonjunior94/bank/loan/infra/environments"
	"github.com/jailtonjunior94/bank/loan/infra/ioc"
	"github.com/jailtonjunior94/bank/loan/presentation/routes"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const defaultPort = "9000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	prometheus := fiberprometheus.New("customer-api")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	environments.NewSettings()

	mongoConnection := database.NewConnection()
	defer mongoConnection.Disconnect()

	ioc.SetupDependencyInjection(mongoConnection)
	routes.SetupRoutes(app)

	fmt.Printf("🚀 Customer API is running on http://localhost:%v", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
