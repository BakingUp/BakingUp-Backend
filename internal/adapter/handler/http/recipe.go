package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type RecipeHandler struct {
	svc port.RecipeService
}

func NewRecipeHandler(svc port.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		svc: svc,
	}
}

func (rh *RecipeHandler) GetAllRecipes(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	recipes, err := rh.svc.GetAllRecipes(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all recipes", err.Error())
	}

	handleSuccess(c, recipes)
	return nil
}
