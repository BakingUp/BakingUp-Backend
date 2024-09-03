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
	}

	handleSuccess(c, stocks)
	return nil
}
