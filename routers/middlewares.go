package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonp17/gosquoosh/config"
)

func APIKeyAuth(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey != config.API_KEY_MICROSERVICE {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}

// Middleware para añadir X-Robots-Tag: noindex
func NoIndexMiddleware(c *fiber.Ctx) error {
	c.Set("X-Robots-Tag", "noindex")
	return c.Next()
}
