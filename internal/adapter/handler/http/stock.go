package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	svc port.StockService
}

func NewStockHandler(svc port.StockService) *StockHandler {
	return &StockHandler{
		svc: svc,
	}
}

func (sh *StockHandler) GetAllStocks(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	stocks, err := sh.svc.GetAllStocks(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all stocks.", err.Error())
		return nil
	}

	handleSuccess(c, stocks)
	return nil
}

// GetStockDetail godoc
// @Summary      Get stock details
// @Description  Get stock details by recipe ID and user ID
// @Tags         stock
// @Accept       json
// @Produce      json
// @Param        recipe_id  query  string  true  "Recipe ID"
// @Success      200  {object}  domain.StockDetail  "Success"
// @Failure      400  {object}  response     "Cannot get stock detail"
// @Router       /stock/getStockDetail [get]
func (sh *StockHandler) GetStockDetail(c *fiber.Ctx) error {
	recipeID := c.Query("recipe_id")

	stock, err := sh.svc.GetStockDetail(c, recipeID)
	if err != nil {
		handleError(c, 400, "Cannot get stock detail.", err.Error())
		return nil
	}

	handleSuccess(c, stock)
	return nil
}

// DeleteStock godoc
// @Summary      Delete a stock
// @Description  Delete a stock by recipe id
// @Tags         stock
// @Accept       json
// @Produce      json
// @Param        recipe_id  query  string  true  "Recipe ID"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response     "Cannot delete a stock"
// @Router       /stock/deleteStock [delete]
func (sh *StockHandler) DeleteStock(c *fiber.Ctx) error {
	recipeID := c.Query("recipe_id")

	err := sh.svc.DeleteStock(c, recipeID)
	if err != nil {
		handleError(c, 400, "Cannot delete a stock.", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}
