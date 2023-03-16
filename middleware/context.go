package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func FiberContextFromMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_ = context.WithValue(c.Context(), "FiberContextKey", c)
		c.Context().SetUserValue("FiberContextKey", c)
		return c.Next()
	}
}
