package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type OrderRepository interface {
	GetAllOrders(c *fiber.Ctx, userID string) ([]db.OrdersModel, error)
	GetOrderDetail(c *fiber.Ctx, orderID string) (*db.OrdersModel, error)
	DeleteOrder(c *fiber.Ctx, orderID string) error
	AddInStoreOrder(c *fiber.Ctx, inStoreOrder *domain.AddInStoreOrderRequest) error
	AddPreOrderOrder(c *fiber.Ctx, preOrderOrder *domain.AddPreOrderOrderRequest) error
	EditOrderStatus(c *fiber.Ctx, orderStatue *domain.EditOrderStatusRequest) error
}

type OrderService interface {
	GetAllOrders(c *fiber.Ctx, userID string) (*domain.Orders, error)
	GetOrderDetail(c *fiber.Ctx, orderID string) (interface{}, error)
	DeleteOrder(c *fiber.Ctx, orderID string) error
	AddInStoreOrder(c *fiber.Ctx, inStoreOrder *domain.AddInStoreOrderRequest) error
	AddPreOrderOrder(c *fiber.Ctx, preOrderOrder *domain.AddPreOrderOrderRequest) error
	EditOrderStatus(c *fiber.Ctx, orderStatue *domain.EditOrderStatusRequest) error
	AddOrderNotification(c *fiber.Ctx, orderProducts []domain.OrderProduct, userId string) error
	EditPreOrderStatusNotification(c *fiber.Ctx, orderID string) error
}
