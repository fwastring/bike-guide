package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/handlers"
	authMiddleware "backend/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth routes
	api.Post("/login", handlers.Login)
	api.Get("/logout", handlers.Logout)

	// Protected group
	routes := api.Group("/routes", authMiddleware.RequireLogin)
	routes.Get("/", handlers.GetRoutes)
	routes.Put("/:id", handlers.UpdateRoute)
	routes.Post("/", handlers.CreateRoute)
	routes.Get("/:id", handlers.GetRoute)
	routes.Delete("/:id", handlers.DeleteRoute)

	// Locations
	location := api.Group("/location", authMiddleware.RequireLogin)
	location.Post("/coords", handlers.GetCoordinates)
	location.Post("/", handlers.FindRoute)
	location.Post("/overpass", handlers.Overpass)
}
