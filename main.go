package main

import (
	apiv0 "ai-audience-analytics/api/v0"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	apiv0.SetupRoutes(app)

	app.Listen(":3000")
}
