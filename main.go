package main

import (
	apiv0 "ai-audience-analytics/api/v0"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(CustomCORS())
	apiv0.SetupRoutes(app)

	appPort := 8080
	appErr := app.Listen(fmt.Sprintf(":%d", appPort))
	if appErr != nil {
		panic(appErr)
	} else {
		println(fmt.Sprintf("Server is running on port %d", appPort))
	}
}
