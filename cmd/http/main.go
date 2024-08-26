package main

import (
	"log/slog"
	"os"

	"github.com/BakingUp/BakingUp-Backend/internal/adapter/config"
	"github.com/BakingUp/BakingUp-Backend/internal/infrastructure"
	"github.com/gofiber/fiber/v2"
)

func main() {
	
	app := fiber.New()
	config, err := config.New()
	
	client := infrastructure.InitializePrismaClient()
    defer client.Disconnect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to BakingUp Backend API")
	})
	
	port := config.HTTP.Port
	if port == "" {
		port = "8000"
	}
	app.Listen(":" + port)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}