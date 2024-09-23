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
)

// @title         BakingUp Backend API
// @version       1.0
// @description   This is the BakingUp Backend API.
// @host          localhost:8000
// @BasePath      /api

func main() {

	app := fiber.New()
	config, err := config.New()

	if err != nil {
		slog.Error("Error loading the configuration", "error", err)
		os.Exit(1)
	}

	client := infrastructure.InitializePrismaClient()
	defer client.Disconnect()

	http.SetupCORS(app, config.HTTP.AllowedOrigins)

	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	authHandler := http.NewAuthHandler(userService)
	userHandler := http.NewUserHandler(userService)

	ingredientRepo := repository.NewIngredientRepository(client)
	ingredientService := service.NewIngredientService(ingredientRepo, userService)
	ingredientHandler := http.NewIngredientHandler(ingredientService)

	recipeRepo := repository.NewRecipeRepository(client)
	recipeService := service.NewRecipeService(recipeRepo, userService)
	recipeHandler := http.NewRecipeHandler(recipeService)

	stockRepo := repository.NewStockRepository(client)
	stockService := service.NewStockService(stockRepo, userService)
	stockHandler := http.NewStockHandler(stockService)

	orderRepo := repository.NewOrderRespository(client)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := http.NewOrderHandler(orderService)

	_, err = http.NewRouter(app, *ingredientHandler, *recipeHandler, *authHandler, *stockHandler, *userHandler, *orderHandler)

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
