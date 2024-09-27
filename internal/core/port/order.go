package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type OrderRespository interface {
	GetAllOrders(c *fiber.Ctx, userID string) ([]db.OrdersModel, error)
}

type OrderService interface {
	GetAllOrders(c *fiber.Ctx, userID string) (*domain.Orders, error)
}