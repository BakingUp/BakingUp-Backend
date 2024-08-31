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

func (ih *IngredientHandler) GetIngredientDetail(c *fiber.Ctx) error {
    ingredientID := c.Params("ingredientID")

    ingredient, err := ih.svc.GetIngredientDetail(c, ingredientID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

	rsp := newIngredientDetailResponse(ingredient)

    handleSuccess(c, rsp)
	return nil
}