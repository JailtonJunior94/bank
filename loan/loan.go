package main

import (
	"fmt"
	"log"
	"os"

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

	fmt.Printf("🚀 Loan API is running on http://localhost:%v", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
