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
	app.Static("/images", "./images")
	configVar, err := config.New()

	if err != nil {
		slog.Error("Error loading the configuration", "error", err)
		os.Exit(1)
	}

	firebaseApp, _, _ := config.SetupFirebase()

	client := infrastructure.InitializePrismaClient()
	defer client.Disconnect()

	http.SetupCORS(app, configVar.HTTP.AllowedOrigins)

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
	stockService := service.NewStockService(stockRepo, userService, ingredientService, recipeRepo)
	stockHandler := http.NewStockHandler(stockService)

	notificationRepo := repository.NewNotificationRepository(client)
	notificationService := service.NewNotificationService(notificationRepo, userService, userRepo, firebaseApp)
	notificationHandler := http.NewNotificationHandler(notificationService)

	orderRepo := repository.NewOrderRespository(client)
	orderService := service.NewOrderService(orderRepo, userRepo, userService, notificationService, stockService, firebaseApp)
	orderHandler := http.NewOrderHandler(orderService)

	settingsRepo := repository.NewSettingsRepository(client)
	settingsService := service.NewSettingsService(settingsRepo, userService)
	settingsHandler := http.NewSetingsHandler(settingsService)

	homeRepo := repository.NewHomeRepository(client)
	homeService := service.NewHomeService(homeRepo, userService, settingsService, settingsRepo, recipeRepo, ingredientRepo)
	homeHandler := http.NewHomeHandler(homeService)

	_, err = http.NewRouter(app, *ingredientHandler, *recipeHandler, *authHandler, *stockHandler, *userHandler, *orderHandler, *settingsHandler, *notificationHandler, *homeHandler)

	port := configVar.HTTP.Port
	if port == "" {
		port = "8000"
	}

	app.Listen(":" + port)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
