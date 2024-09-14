package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type IngredientHandler struct {
	svc port.IngredientService
}

func NewIngredientHandler(svc port.IngredientService) *IngredientHandler {
	return &IngredientHandler{
		svc: svc,
	}
}

func (ih *IngredientHandler) GetAllIngredients(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	ingredients, err := ih.svc.GetAllIngredients(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all ingredients", err.Error())
	}

	handleSuccess(c, ingredients)
	return nil
}

// GetIngredientDetail godoc
// @Summary      Get ingredient details
// @Description  Get ingredient details by ingredient ID
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_id  query  string  true  "Ingredient ID"
// @Success      200  {object}  IngredientDetail  "Success"
// @Failure      400  {object}  response     "Cannot get ingredient detail"
// @Router       /ingredient [get]
func (ih *IngredientHandler) GetIngredientDetail(c *fiber.Ctx) error {
	ingredientID := c.Query("ingredient_id")

	ingredient, err := ih.svc.GetIngredientDetail(c, ingredientID)
	if err != nil {
		handleError(c, 400, "Cannot get ingredient detail", err.Error())
		return nil
	}

	rsp := newIngredientDetailResponse(ingredient)

	handleSuccess(c, rsp)
	return nil
}
