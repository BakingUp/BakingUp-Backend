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
// @Success      200  {object}  domain.IngredientDetail  "Success"
// @Failure      400  {object}  response     "Cannot get ingredient detail"
// @Router       /ingredient/getIngredientDetail [get]
func (ih *IngredientHandler) GetIngredientDetail(c *fiber.Ctx) error {
	ingredientID := c.Query("ingredient_id")

	ingredient, err := ih.svc.GetIngredientDetail(c, ingredientID)
	if err != nil {
		handleError(c, 400, "Cannot get ingredient detail", err.Error())
		return nil
	}

	handleSuccess(c, ingredient)
	return nil
}

// GetIngredientStockDetail godoc
// @Summary      Get ingredient stock details
// @Description  Get ingredient stock details by ingredient stock ID
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_stock_id  query  string  true  "Ingredient stock ID"
// @Success      200  {object}  domain.IngredientStockDetail  "Success"
// @Failure      400  {object}  response     "Cannot get ingredient stock detail"
// @Router       /ingredient/getIngredientStockDetail [get]
func (ih *IngredientHandler) GetIngredientStockDetail(c *fiber.Ctx) error {
	ingredientStockID := c.Query("ingredient_stock_id")

	ingredient, err := ih.svc.GetIngredientStockDetail(c, ingredientStockID)
	if err != nil {
		handleError(c, 400, "Cannot get ingredient stock detail", err.Error())
		return nil
	}

	handleSuccess(c, ingredient)
	return nil
}
