package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stgonzales/go-api-jwt/router"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
