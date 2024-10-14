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

// GetAllStocks godoc
// @Summary      Get all stocks
// @Description  Get all stocks by user ID
// @Tags         stock
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.StockList  "Success"
// @Failure      400  {object}  response     "Cannot get all stocks"
// @Router       /stock/getAllStocks [get]
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

func (sh *StockHandler) GetAllStocksForOrder(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	stocks, err := sh.svc.GetAllStocksForOrder(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all stocks for order page.", err.Error())
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

// GetStockBatch godoc
// @Summary      Get stock batch
// @Description  Get stock batch by stock detail ID
// @Tags         stock
// @Accept       json
// @Produce      json
// @Param        stock_detail_id  query  string  true  "Stock Detail ID"
// @Success      200  {object}  domain.StockBatch  "Success"
// @Failure      400  {object}  response     "Cannot get stock batch"
// @Router       /stock/getStockBatch [get]
func (sh *StockHandler) GetStockBatch(c *fiber.Ctx) error {
	stockDetailID := c.Query("stock_detail_id")

	batch, err := sh.svc.GetStockBatch(c, stockDetailID)
	if err != nil {
		handleError(c, 400, "Cannot get stock batch.", err.Error())
		return nil
	}

	handleSuccess(c, batch)
	return nil
}
