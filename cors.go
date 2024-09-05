package main

import "github.com/gofiber/fiber/v2"

// CustomCORS is a middleware that sets the Access-Control-Allow-Origin header based on the request's Origin header.
func CustomCORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		origin := c.Get("Origin")
		if origin != "" {
			c.Set("Access-Control-Allow-Origin", origin)
			c.Set("Access-Control-Allow-Methods", "POST")
			c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		}
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	}
}
