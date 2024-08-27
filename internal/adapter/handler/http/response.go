package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Status int   `json:"status" example:"200"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(status int, message string, data any) response {
	return response{
		Status: status,
		Message: message,
		Data:    data,
	}
}

type IngredientDetail struct {
    IngredientName     string   `json:"ingredient_name"`
    IngredientQuantity string   `json:"ingredient_quantity"`
    StockAmount        int      `json:"stock_amount"`
    IngredientURL      []string `json:"ingredient_url"`
    IngredientLessThan int      `json:"ingredient_less_than"`
    Stocks             []domain.Stock  `json:"stocks"`
}

func newIngredientDetailResponse(ingredient *domain.IngredientDetail) *IngredientDetail {
	return &IngredientDetail{
		IngredientName:     ingredient.IngredientName,
		IngredientQuantity: ingredient.IngredientQuantity,
		StockAmount:        ingredient.StockAmount,
		IngredientURL:      ingredient.IngredientURL,
		IngredientLessThan: ingredient.IngredientLessThan,
		Stocks:             ingredient.Stocks,
	}
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *fiber.Ctx, data any) {
	rsp := newResponse(200, "Success", data)
	ctx.JSON(rsp)
}