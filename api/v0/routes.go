package apiv0

import (
	events "ai-audience-analytics/internal"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/v0/collect", func(c *fiber.Ctx) error {
		queryParams := c.Queries()
		events.WriteEvent(queryParams)
		return c.SendStatus(fiber.StatusOK)
	})
}
