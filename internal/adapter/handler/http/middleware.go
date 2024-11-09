package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetupCORS(app *fiber.App, allowedOrigins string) {
    app.Use(func(c *fiber.Ctx) error {
        c.Set("Access-Control-Allow-Origin", allowedOrigins)
        c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        return c.Next()
    })
}