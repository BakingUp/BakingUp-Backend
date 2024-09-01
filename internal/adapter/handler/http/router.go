package http

import "github.com/gofiber/fiber/v2"

type Router struct {
	router fiber.Router
}

func NewRouter(a *fiber.App, ingredientHandler IngredientHandler, recipeHandler RecipeHandler) (*Router, error) {

	api := a.Group("/api")
	{
		ingredient := api.Group("/ingredient")
		{
			ingredient.Get("/getAllIngredients", ingredientHandler.GetAllIngredients)
			ingredient.Get("/:ingredientID", ingredientHandler.GetIngredientDetail)
		}

		recipe := api.Group("/recipe")
		{
			recipe.Get("/getAllRecipes", recipeHandler.GetAllRecipes)
		}
	}

	return &Router{api}, nil
}
