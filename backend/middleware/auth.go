package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireLogin(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get session",
		})
	}

	userID := sess.Get("userId")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You must be logged in",
		})
	}

	// Optionally attach user info to context
	c.Locals("userId", userID)
	return c.Next()
}

