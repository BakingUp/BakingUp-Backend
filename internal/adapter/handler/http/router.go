package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	router fiber.Router
}

func NewRouter(a *fiber.App, ingredientHandler IngredientHandler, recipeHandler RecipeHandler, authHandler AuthHandler, stockHandler StockHandler, userHandler UserHandler) (*Router, error) {
	a.Get("/swagger/*", swagger.HandlerDefault)

	api := a.Group("/api")
	{
		user := api.Group("/user")
		{
			user.Get("/getUserInfo", userHandler.GetUserInfo)
		}
		auth := api.Group("/auth")
		{
			auth.Post("/register", authHandler.Register)
			auth.Post("/addDeviceToken", authHandler.AddDeviceToken)
			auth.Delete("/deleteDeviceToken", authHandler.DeleteDeviceToken)
			auth.Delete("/deleteAllExceptDeviceToken", authHandler.DeleteAllExceptDeviceToken)
		}
		ingredient := api.Group("/ingredient")
		{
			ingredient.Get("/getAllIngredients", ingredientHandler.GetAllIngredients)
			ingredient.Get("/getIngredientDetail", ingredientHandler.GetIngredientDetail)
			ingredient.Get("/getIngredientStockDetail", ingredientHandler.GetIngredientStockDetail)
		}

		recipe := api.Group("/recipe")
		{
			recipe.Get("/getAllRecipes", recipeHandler.GetAllRecipes)
		}

		stock := api.Group("/stock")
		{
			stock.Get("/getAllStocks", stockHandler.GetAllStocks)
		}
	}

	return &Router{api}, nil
}
