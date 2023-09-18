package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stgonzales/go-api-jwt/handler"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		logger.New()
		return c.Next()
	})

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
}
