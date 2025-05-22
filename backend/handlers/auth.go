package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/middleware"
)

func Login(c *fiber.Ctx) error {
	sess, err := middleware.Store.Get(c)
	if err != nil {
		return err
	}
	sess.Set("userId", "some-user-id") // In real life, after validation
	sess.Save()
	return c.JSON(fiber.Map{"message": "Logged in"})
}

func ProtectedRoute(c *fiber.Ctx) error {
	sess, err := middleware.Store.Get(c)
	if err != nil {
		return err
	}
	if sess.Get("userId") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	return c.JSON(fiber.Map{"message": "You're authorized!"})
}

func Logout(c *fiber.Ctx) error {
	sess, err := middleware.Store.Get(c)
	if err != nil {
		return err
	}
	sess.Destroy()
	return c.JSON(fiber.Map{"message": "Logged out"})
}
