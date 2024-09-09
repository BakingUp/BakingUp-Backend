package main

import (
	"log/slog"
	"os"

	_ "github.com/BakingUp/BakingUp-Backend/docs"
	"github.com/BakingUp/BakingUp-Backend/internal/adapter/config"
	"github.com/BakingUp/BakingUp-Backend/internal/adapter/handler/http"
	"github.com/BakingUp/BakingUp-Backend/internal/adapter/storage/postgres/repository"
	"github.com/BakingUp/BakingUp-Backend/internal/core/service"
	"github.com/BakingUp/BakingUp-Backend/internal/infrastructure"
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

	if err != nil {
		slog.Error("Error loading the configuration", "error", err)
		os.Exit(1)
	}

	client := infrastructure.InitializePrismaClient()
	defer client.Disconnect()

	http.SetupSwagger(app)
	http.SetupCORS(app, config.HTTP.AllowedOrigins)

	app.Get("/", welcome)

	app.Get("/swagger/*", swagger.HandlerDefault)

	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	authHandler := http.NewAuthHandler(userService)

	ingredientRepo := repository.NewIngredientRepository(client)
	ingredientService := service.NewIngredientService(ingredientRepo, userService)
	ingredientHandler := http.NewIngredientHandler(ingredientService)

	recipeRepo := repository.NewRecipeRepository(client)
	recipeService := service.NewRecipeService(recipeRepo, userService)
	recipeHandler := http.NewRecipeHandler(recipeService)

	stockRepo := repository.NewStockRepository(client)
	stockService := service.NewStockService(stockRepo, userService)
	stockHandler := http.NewStockHandler(stockService)

	_, err = http.NewRouter(app, *ingredientHandler, *recipeHandler, *authHandler, *stockHandler)

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
