package main

import (
	"log/slog"
	"os"

	"github.com/BakingUp/BakingUp-Backend/internal/adapter/config"
	"github.com/BakingUp/BakingUp-Backend/internal/infrastructure"
	"github.com/BakingUp/BakingUp-Backend/internal/adapter/handler/http"
	_ "github.com/BakingUp/BakingUp-Backend/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title BakingUp Backend API
// @version 1.0
// @description This is the BakingUp Backend API.
// @host localhost:8000
// @BasePath /

// @Summary Welcome message
// @Description Returns a welcome message
// @Tags root
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Welcome to BakingUp Backend API"
// @Router / [get]
func welcome(c *fiber.Ctx) error {
    return c.SendString("Welcome to BakingUp Backend API")
}

func main() {
	
	app := fiber.New()
	config, err := config.New()
	
	client := infrastructure.InitializePrismaClient()
    defer client.Disconnect()

	http.SetupSwagger(app)
	
	app.Get("/", welcome)

	app.Get("/swagger/*", swagger.HandlerDefault)
	
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