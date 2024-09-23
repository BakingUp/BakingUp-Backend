package service

import (
	"fmt"
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type OrderService struct {
	orderRepo port.OrderRespository
}

func NewOrderService(orderRepo port.OrderRespository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) GetAllOrders(c *fiber.Ctx, userID string) (*domain.Orders, error) {
	orders, err := s.orderRepo.GetAllOrders(c, userID)

	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("No orders found for user ID %s", userID)
	}

	var list []domain.OrderDetail
	for _, order := range orders {
		totalPrice := 0.0
		for _, product := range order.OrderProducts() {
			stock, err := product.Recipe().Stocks()
			if err == true {
				totalPrice = totalPrice + (stock.SellingPrice * float64(product.ProductQuantity))
			}
		}

		list = append(list, domain.OrderDetail{
			OrderID:     order.OrderID,
			OrderIndex:  order.OrderIndex,
			Total:       totalPrice,
			OrderDate:   order.OrderDate,
			PickUpDate:  order.PickUpDateTime,
			OrderStatus: order.OrderStatus,
		})
	}

	reponse := &domain.Orders{
		Orders: list,
	}

	return reponse, err

}
