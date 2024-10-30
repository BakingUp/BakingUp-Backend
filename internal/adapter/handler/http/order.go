package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

func (oh *OrderHandler) AddInStoreOrder(c *fiber.Ctx) error {
	var inStoreOrder domain.AddInStoreOrderRequest

	if err := c.BodyParser(&inStoreOrder); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}
	if inStoreOrder.UserID == "" {
		handleError(c, 400, "UserID is required", "")
		return nil
	}

	err := oh.svc.AddInStoreOrder(c, &inStoreOrder)
	if err != nil {
		handleError(c, 400, "Cannot add a new in-store order.", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully add a new in-store order.")
	return nil
}

func (oh *OrderHandler) AddPreOrderOrder(c *fiber.Ctx) error {
	var preOrderOrder domain.AddPreOrderOrderRequest

	if err := c.BodyParser(&preOrderOrder); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}
	if preOrderOrder.UserID == "" {
		handleError(c, 400, "UserID is required", "")
		return nil
	}

	err := oh.svc.AddPreOrderOrder(c, &preOrderOrder)
	if err != nil {
		handleError(c, 400, "Cannot add a new pre-order order.", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully add a new pre-order order.")
	return nil
}

func (oh *OrderHandler) EditOrderStatus(c *fiber.Ctx) error {
	var orderStatue domain.EditOrderStatusRequest

	if err := c.BodyParser(&orderStatue); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}
	if orderStatue.OrderID == "" {
		handleError(c, 400, "OrderID is required", "")
		return nil
	}

	err := oh.svc.EditOrderStatus(c, &orderStatue)
	if err != nil {
		handleError(c, 400, "Cannot edit an order status.", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully edit an order status.")
	return nil
}
