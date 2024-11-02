package service

import (
	"fmt"
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type OrderService struct {
	orderRepo   port.OrderRepository
	userService port.UserService
}

func NewOrderService(orderRepo port.OrderRepository, userService port.UserService) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		userService: userService,
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

	var list []domain.OrderInfo
	for _, order := range orders {
		totalPrice := 0.0
		for _, product := range order.OrderProducts() {
			stock, err := product.Recipe().Stocks()
			if err == true {
				totalPrice = totalPrice + (stock.SellingPrice * float64(product.ProductQuantity))
			}
		}
		if order.IsPreOrder {
			//Pre-Order case
			list = append(list, domain.OrderInfo{
				OrderID:       order.OrderID,
				OrderIndex:    order.OrderIndex,
				OrderPlatform: order.OrderPlatform,
				IsPreOrder:    order.IsPreOrder,
				Total:         totalPrice,
				OrderDate:     order.OrderDate.Format("02/01/2006 03:04 PM"),
				PickUpDate:    order.PickUpDateTime.Format("02/01/2006 03:04 PM"),
				OrderStatus:   order.OrderStatus,
			})
		} else {
			// In-store case
			list = append(list, domain.OrderInfo{
				OrderID:       order.OrderID,
				OrderIndex:    order.OrderIndex,
				OrderPlatform: order.OrderPlatform,
				IsPreOrder:    order.IsPreOrder,
				Total:         totalPrice,
				OrderDate:     order.OrderDate.Format("02/01/2006 03:04 PM"),
				OrderStatus:   order.OrderStatus,
			})
		}

	}

	reponse := &domain.Orders{
		Orders: list,
	}

	return reponse, err

}

func (s *OrderService) GetOrderDetail(c *fiber.Ctx, orderID string) (interface{}, error) {
	order, err := s.orderRepo.GetOrderDetail(c, orderID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, order.UserID)
	if err != nil {
		return nil, err
	}

	var list []domain.OrderStock
	totalPrice := 0.0
	cost := 0.0

	for _, product := range order.OrderProducts() {
		recipe := product.Recipe()
		var recipeImg string
		if recipe != nil {
			for _, recipeImageItem := range recipe.RecipeImages() {
				if recipeImageItem.RecipeID == recipe.RecipeID {
					recipeImg = recipeImageItem.RecipeURL
				}
			}
		}
		stock, err := product.Recipe().Stocks()
		if err == true {
			totalPrice = totalPrice + (stock.SellingPrice * float64(product.ProductQuantity))
			cost = cost + (stock.Cost * float64(product.ProductQuantity))
			list = append(list, domain.OrderStock{
				Name:       util.GetRecipeName(product.Recipe(), language),
				Quantity:   product.ProductQuantity,
				StockPrice: stock.SellingPrice,
				ImgURL:     recipeImg,
			})
		}
	}
	profit := totalPrice - cost

	if order.IsPreOrder {
		// Pre-order case
		detail := &domain.PreOrderOrderDetails{
			OrderIndex:    order.OrderIndex,
			OrderStatus:   order.OrderStatus,
			OrderPlatform: order.OrderPlatform,
			OrderDate:     order.OrderDate.Format("02/01/2006"),
			OrderTime:     order.OrderDate.Format("03:04 PM"),
			OrderType:     order.OrderType,
			PickUpDate:    order.PickUpDateTime.Format("02/01/2006"),
			PickUpTime:    order.PickUpDateTime.Format("03:04 PM"),
			OrderTakenBy:  order.OrderTakenBy,
			OrderStock:    list,
			ToTal:         totalPrice,
			Profit:        profit,
		}

		custumerName, ok := order.CustomerName()
		if ok {
			detail.CustomerName = custumerName
		}
		phoneNumber, ok := order.CustomerPhoneNum()
		if ok {
			detail.PhoneNumber = phoneNumber
		}
		pickUpMethod, ok := order.PickUpMethod()
		if ok {
			detail.PickUpMethod = pickUpMethod
		}
		orderNoteText, ok1 := order.OrderNoteText()
		date, ok2 := order.OrderNoteCreateAt()
		if ok1 && ok2 {
			detail.OrderNoteText = orderNoteText
			detail.OrderNoteCreateAt = date.Format("02/01/2006")
		}

		return detail, nil
	} else {
		detail := &domain.InStoreOrderDetails{
			OrderIndex:    order.OrderIndex,
			OrderStatus:   order.OrderStatus,
			OrderPlatform: order.OrderPlatform,
			OrderDate:     order.OrderDate.Format("02/01/2006"),
			OrderTime:     order.OrderDate.Format("03:04 PM"),
			OrderType:     order.OrderType,
			OrderTakenBy:  order.OrderTakenBy,
			OrderStock:    list,
			ToTal:         totalPrice,
			Profit:        profit,
		}
		orderNoteText, ok1 := order.OrderNoteText()
		date, ok2 := order.OrderNoteCreateAt()
		if ok1 && ok2 {
			detail.OrderNoteText = orderNoteText
			detail.OrderNoteCreateAt = date.Format("02/01/2006")
		}

		return detail, nil
	}
}

func (s *OrderService) DeleteOrder(c *fiber.Ctx, orderID string) error {
	err := s.orderRepo.DeleteOrder(c, orderID)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) AddInStoreOrder(c *fiber.Ctx, inStoreOrder *domain.AddInStoreOrderRequest) error {
	err := s.orderRepo.AddInStoreOrder(c, inStoreOrder)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) AddPreOrderOrder(c *fiber.Ctx, preOrderOrder *domain.AddPreOrderOrderRequest) error {
	err := s.orderRepo.AddPreOrderOrder(c, preOrderOrder)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) EditOrderStatus(c *fiber.Ctx, orderStatue *domain.EditOrderStatusRequest) error {
	err := s.orderRepo.EditOrderStatus(c, orderStatue)
	if err != nil {
		return err
	}
	return nil
}
