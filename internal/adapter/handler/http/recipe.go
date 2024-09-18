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

// GetRecipeDetail godoc
// @Summary      Get recipe details
// @Description  Get recipe details by recipe ID
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        recipe_id  query  string  true  "Recipe ID"
// @Success      200  {object}  domain.RecipeDetail  "Success"
// @Failure      400  {object}  response     "Cannot get recipe detail"
// @Router       /recipe/getRecipeDetail [get]
func (rh *RecipeHandler) GetRecipeDetail(c *fiber.Ctx) error {
	recipeID := c.Query("recipe_id")

	recipe, err := rh.svc.GetRecipeDetail(c, recipeID)
	if err != nil {
		handleError(c, 400, "Cannot get recipe detail", err.Error())
		return nil
	}

	handleSuccess(c, recipe)
	return nil
}
