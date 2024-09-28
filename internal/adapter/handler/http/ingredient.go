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

// GetAllIngredients godoc
// @Summary      Get all ingredients
// @Description  Get all ingredients by user ID
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.IngredientList  "Success"
// @Failure      400  {object}  response     "Cannot get all ingredients"
// @Router       /ingredient/getAllIngredients [get]
func (ih *IngredientHandler) GetAllIngredients(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	ingredients, err := ih.svc.GetAllIngredients(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all ingredients", err.Error())
		return nil
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

// DeleteIngredientBatchNote godoc
// @Summary      Delete ingredient batch note
// @Description  Delete ingredient batch note by ingredient note ID
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_note_id  query  string  true  "Ingredient Note ID"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot delete ingredient batch note"
// @Router       /ingredient/deleteIngredientBatchNote [delete]
func (ih *IngredientHandler) DeleteIngredientBatchNote(c *fiber.Ctx) error {
	ingredientNoteID := c.Query("ingredient_note_id")

	err := ih.svc.DeleteIngredientBatchNote(c, ingredientNoteID)
	if err != nil {
		handleError(c, 400, "Cannot delete ingredient batch note", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}

// DeleteIngredient godoc
// @Summary      Delete an ingredient
// @Description  Delete an ingredient by using ingredient id
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_id  query  string  true  "Ingredient ID"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot delete an ingredient"
// @Router       /ingredient/deleteIngredient [delete]
func (ih *IngredientHandler) DeleteIngredient(c *fiber.Ctx) error {
	ingredientID := c.Query("ingredient_id")

	err := ih.svc.DeleteIngredient(c, ingredientID)
	if err != nil {
		handleError(c, 400, "Cannot delete an ingredient", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}
