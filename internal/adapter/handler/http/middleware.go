package http

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func SetupSwagger(app *fiber.App) {
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

    app.Use(swagger.New(cfg))
}

func SetupCORS(app *fiber.App, allowedOrigins string) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", allowedOrigins)
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return c.Next()
	})
}