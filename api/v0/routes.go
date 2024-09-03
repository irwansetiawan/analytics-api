package apiv0

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/v0/collect", func(c *fiber.Ctx) error {
		queryParams := c.Queries()
		return c.JSON(queryParams)
	})
}
