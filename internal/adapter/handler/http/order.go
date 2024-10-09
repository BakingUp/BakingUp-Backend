package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	svc port.OrderService
}

func NewOrderHandler(svc port.OrderService) *OrderHandler {
	return &OrderHandler{
		svc: svc,
	}
}

func (oh *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	orders, err := oh.svc.GetAllOrders(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all orders.", err.Error())
		return nil
	}

	handleSuccess(c, orders)
	return nil
}

func (oh *OrderHandler) GetOrderDeatil(c *fiber.Ctx) error {
	orderID := c.Query("order_id")

	orderDetail, err := oh.svc.GetOrderDetail(c, orderID)
	if err != nil {
		handleError(c, 400, "Cannot get an order detail", err.Error())
		return nil
	}

	handleSuccess(c, orderDetail)
	return nil
}

func (oh *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Query("order_id")

	err := oh.svc.DeleteOrder(c, orderID)
	if err != nil {
		handleError(c, 400, "Cannot delete an order.", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}
