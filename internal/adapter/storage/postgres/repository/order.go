package repository

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type OrderRespository struct {
	db *db.PrismaClient
}

func NewOrderRespository(db *db.PrismaClient) *OrderRespository {
	return &OrderRespository{
		db: db,
	}
}

func (or *OrderRespository) GetAllOrders(c *fiber.Ctx, userID string) ([]db.OrdersModel, error) {
	orders, err := or.db.Orders.FindMany(
		db.Orders.UserID.Equals(userID)).With(db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch()))).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (or *OrderRespository) GetOrderDetail(c *fiber.Ctx, orderID string) (*db.OrdersModel, error) {
	order, err := or.db.Orders.FindFirst(
		db.Orders.OrderID.Equals(orderID)).With(db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch()))).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return order, nil
}
