package http

import "github.com/gofiber/fiber/v2"

type Router struct {
	router fiber.Router
}

func NewRouter(a *fiber.App, ingredientHandler IngredientHandler, recipeHandler RecipeHandler, authHandler AuthHandler, stockHandler StockHandler) (*Router, error) {

	api := a.Group("/api")
	{
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
			ingredient.Get("/:ingredientID", ingredientHandler.GetIngredientDetail)
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
