package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

// GetAddEditIngredientStockDetail godoc
// @Summary      Get add edit ingredient stock details
// @Description  Get add edit ingredient stock details by ingredient ID
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_id  query  string  true  "Ingredient ID"
// @Success      200  {object}  domain.AddEditIngredientStockDetail  "Success"
// @Failure      400  {object}  response     "Cannot get add edit ingredient stock detail"
// @Router       /ingredient/getAddEditIngredientStockDetail [get]
func (ih *IngredientHandler) GetAddEditIngredientStockDetail(c *fiber.Ctx) error {
	ingredientID := c.Query("ingredient_id")

	ingredient, err := ih.svc.GetAddEditIngredientStockDetail(c, ingredientID)

	if err != nil {
		handleError(c, 400, "Cannot get add edit ingredient stock detail", err.Error())
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

// DeleteIngredientStock godoc
// @Summary      Delete an ingredient stock
// @Description  Delete an ingredient stock by using ingredient stock id
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_stock_id  query  string  true  "Ingredient Stock ID"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot delete an ingredient stock"
// @Router       /ingredient/deleteIngredientStock [delete]
func (ih *IngredientHandler) DeleteIngredientStock(c *fiber.Ctx) error {
	ingredientStockID := c.Query("ingredient_stock_id")

	err := ih.svc.DeleteIngredientStock(c, ingredientStockID)
	if err != nil {
		handleError(c, 400, "Cannot delete an ingredient stock", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}

// AddIngredient godoc
// @Summary      Add ingredient
// @Description  Add ingredient by using ingredient request
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        AddIngredientRequest  body  domain.AddIngredientRequest  true  "Ingredient Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot add ingredients"
// @Router       /ingredient/addIngredients [post]
func (ih *IngredientHandler) AddIngredient(c *fiber.Ctx) error {
    var addIngredientRequest domain.AddIngredientRequest
    if err := c.BodyParser(&addIngredientRequest); err != nil {
        handleError(c, 400, "Cannot add ingredients", err.Error())
        return nil
    }
	
    err := ih.svc.AddIngredient(c, &addIngredientRequest)
    if err != nil {
        handleError(c, 400, "Cannot add ingredients", err.Error())
        return nil
    }

    handleSuccess(c, nil)
    return nil
}

// AddIngredientStock godoc
// @Summary      Add ingredient stock
// @Description  Add ingredient stock by using ingredient stock request
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        AddIngredientStockRequest  body  domain.AddIngredientStockRequest  true  "Ingredient Stock Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot add ingredient stock"'
// @Router       /ingredient/addIngredientStock [post]
func (ih *IngredientHandler) AddIngredientStock(c *fiber.Ctx) error {
	var addIngredientStock domain.AddIngredientStockRequest

	if err := c.BodyParser(&addIngredientStock); err != nil {
		handleError(c, 400, "Cannot add ingredient stock", err.Error())
		return nil
	}

	err := ih.svc.AddIngredientStock(c, &addIngredientStock)
	if err != nil {
		handleError(c, 400, "Cannot add ingredient stock", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully add ingredient stock")
	return nil
}

// EditIngredient godoc
// @Summary      Edit ingredient
// @Description  Edit ingredient by using edit ingredient request
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        EditIngredientRequest  body  domain.EditIngredientRequest  true  "Edit Ingredient Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot edit ingredient"
// @Router       /ingredient/editIngredient [put]
func (ih *IngredientHandler) EditIngredient(c *fiber.Ctx) error {
	var editIngredientRequest domain.EditIngredientRequest

	if err := c.BodyParser(&editIngredientRequest); err != nil {
		handleError(c, 400, "Cannot edit ingredient", err.Error())
		return nil
	}

	err := ih.svc.EditIngredient(c, &editIngredientRequest)
	if err != nil {
		handleError(c, 400, "Cannot edit ingredient", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully edit ingredient")
	return nil
}

// GetAddEditIngredientDetail godoc
// @Summary      Get add edit ingredient detail
// @Description  Get add edit ingredient detail by ingredient ID
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_id  query  string  true  "Ingredient ID"
// @Success      200  {object}  domain.GetAddEditIngredientDetail  "Success"
// @Failure      400  {object}  response     "Cannot get add edit ingredient detail"
// @Router       /ingredient/getAddEditIngredientDetail [get]
func (ih *IngredientHandler) GetAddEditIngredientDetail(c *fiber.Ctx) error {
	ingredientID := c.Query("ingredient_id")

	ingredient, err := ih.svc.GetAddEditIngredientDetail(c, ingredientID)
	if err != nil {
		handleError(c, 400, "Cannot get add edit ingredient detail", err.Error())
		return nil
	}

	handleSuccess(c, ingredient)
	return nil
}

// EditIngredientStock godoc
// @Summary      Edit ingredient stock
// @Description  Edit ingredient stock by using edit ingredient stock request
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        EditIngredientStockRequest  body  domain.EditIngredientStockRequest  true  "Edit Ingredient Stock Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot edit ingredient stock"
// @Router       /ingredient/editIngredientStock [put]
func (ih *IngredientHandler) EditIngredientStock(c *fiber.Ctx) error {
	var editIngredientStockRequest domain.EditIngredientStockRequest

	if err := c.BodyParser(&editIngredientStockRequest); err != nil {
		handleError(c, 400, "Cannot edit ingredient stock", err.Error())
		return nil
	}

	err := ih.svc.EditIngredientStock(c, &editIngredientStockRequest)
	if err != nil {
		handleError(c, 400, "Cannot edit ingredient stock", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully edit ingredient stock")
	return nil
}

// GetEditIngredientStockDetail godoc
// @Summary      Get edit ingredient stock detail
// @Description  Get edit ingredient stock detail by ingredient stock ID
// @Tags         ingredient
// @Accept       json
// @Produce      json
// @Param        ingredient_stock_id  query  string  true  "Ingredient Stock ID"
// @Success      200  {object}  domain.GetEditIngredientStockDetail  "Success"
// @Failure      400  {object}  response     "Cannot get edit ingredient stock detail"
// @Router       /ingredient/getEditIngredientStockDetail [get]
func (ih *IngredientHandler) GetEditIngredientStockDetail(c *fiber.Ctx) error {
	ingredientStockID := c.Query("ingredient_stock_id")

	ingredient, err := ih.svc.GetEditIngredientStockDetail(c, ingredientStockID)
	if err != nil {
		handleError(c, 400, "Cannot get edit ingredient stock detail", err.Error())
		return nil
	}

	handleSuccess(c, ingredient)
	return nil
}