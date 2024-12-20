package service

import (
	"fmt"
	"math"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/BakingUp/BakingUp-Backend/internal/adapter/config"
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type OrderService struct {
	orderRepo    port.OrderRepository
	userRepo     port.UserRepository
	userService  port.UserService
	notiService  port.NotificationService
	stockService port.StockService
	firebaseApp  *firebase.App
}

func NewOrderService(orderRepo port.OrderRepository, userRepo port.UserRepository, userService port.UserService, notiService port.NotificationService, stockService port.StockService, firebaseApp *firebase.App) *OrderService {
	return &OrderService{
		orderRepo:    orderRepo,
		userRepo:     userRepo,
		userService:  userService,
		notiService:  notiService,
		stockService: stockService,
		firebaseApp:  firebaseApp,
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

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, err
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

		orderDateThai := order.OrderDate.In(location)
		pickUpDateThai := order.PickUpDateTime.In(location)

		if order.IsPreOrder {
			//Pre-Order case
			list = append(list, domain.OrderInfo{
				OrderID:       order.OrderID,
				OrderIndex:    order.OrderIndex,
				OrderPlatform: order.OrderPlatform,
				IsPreOrder:    order.IsPreOrder,
				Total:         totalPrice,
				OrderDate:     orderDateThai.Format("02/01/2006 03:04 PM"),
				PickUpDate:    pickUpDateThai.Format("02/01/2006 03:04 PM"),
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
				OrderDate:     orderDateThai.Format("02/01/2006 03:04 PM"),
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

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, err
	}

	orderDateThai := order.OrderDate.In(location)
	pickUpDateThai := order.PickUpDateTime.In(location)

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
			OrderDate:     orderDateThai.Format("02/01/2006"),
			OrderTime:     orderDateThai.Format("03:04 PM"),
			OrderType:     order.OrderType,
			PickUpDate:    pickUpDateThai.Format("02/01/2006"),
			PickUpTime:    pickUpDateThai.Format("03:04 PM"),
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
			detail.OrderNoteCreateAt = date.In(location).Format("02/01/2006")
		}

		return detail, nil
	} else {
		detail := &domain.InStoreOrderDetails{
			OrderIndex:    order.OrderIndex,
			OrderStatus:   order.OrderStatus,
			OrderPlatform: order.OrderPlatform,
			OrderDate:     orderDateThai.Format("02/01/2006"),
			OrderTime:     orderDateThai.Format("03:04 PM"),
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
			detail.OrderNoteCreateAt = date.In(location).Format("02/01/2006")
		}

		return detail, nil
	}
}

func (s *OrderService) GetPreOrderOrderDetail(c *fiber.Ctx, orderID string) (*domain.PreOrderOrderDetailNotification, error) {
	order, err := s.orderRepo.GetOrderDetail(c, orderID)
	if err != nil {
		return nil, err
	}

	var orderProductList []domain.OrderProduct

	for _, product := range order.OrderProducts() {

		orderProductList = append(orderProductList, domain.OrderProduct{
			RecipeID:        product.RecipeID,
			ProductQuantity: product.ProductQuantity,
		})
	}

	response := &domain.PreOrderOrderDetailNotification{
		UserID:       order.UserID,
		OrderProduct: orderProductList,
	}

	return response, nil

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

	err = s.AddOrderNotification(c, inStoreOrder.OrderProducts, inStoreOrder.UserID)
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

	status := strings.ToLower(preOrderOrder.OrderStatus)
	if status == "done" {
		err = s.AddOrderNotification(c, preOrderOrder.OrderProducts, preOrderOrder.UserID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *OrderService) EditOrderStatus(c *fiber.Ctx, orderStatue *domain.EditOrderStatusRequest) error {
	err := s.orderRepo.EditOrderStatus(c, orderStatue)
	if err != nil {
		return err
	}

	status := strings.ToLower(orderStatue.OrderStatus)
	if status == "done" {
		err = s.EditPreOrderStatusNotification(c, orderStatue.OrderID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) AddOrderNotification(c *fiber.Ctx, orderProducts []domain.OrderProduct, userId string) error {
	stocks, err := s.stockService.GetAllStocks(c, userId)
	if err != nil {
		return err
	}

	recipeMap := make(map[string]string)

	for _, item := range orderProducts {
		recipeMap[item.RecipeID] = item.RecipeID
	}

	for _, stock := range stocks.Stocks {

		if stock.StockId == recipeMap[stock.StockId] && stock.Quantity < stock.StockLessThan {

			deviceToken, err := s.userRepo.GetDeviceToken(userId)
			if err != nil {
				return err
			}

			if stock.Quantity == 0 {
				err = config.SendToToken(
					s.firebaseApp,
					*deviceToken,
					"Sold out!",
					fmt.Sprintf("We're currently sold out of %s.", stock.StockName),
				)
				if err != nil {
					return err
				}

				err = s.notiService.CreateNotification(c, &domain.CreateNotificationItem{
					UserID:       userId,
					EngTitle:     "Sold out!",
					EngMessage:   fmt.Sprintf("We're currently sold out of %s.", stock.StockName),
					IsRead:       false,
					NotiType:     "ALERT",
					ItemID:       stock.StockId,
					ItemName:     stock.StockName,
					NotiItemType: "STOCK",
				})
				if err != nil {
					return err
				}
			} else {
				err = config.SendToToken(
					s.firebaseApp,
					*deviceToken,
					"Stock Up Time!",
					fmt.Sprintf("%s is running low. Only %d left in stock", stock.StockName, stock.Quantity),
				)
				if err != nil {
					return err
				}
				err = s.notiService.CreateNotification(c, &domain.CreateNotificationItem{
					UserID:       userId,
					EngTitle:     "Stock Up Time!",
					EngMessage:   fmt.Sprintf("%s is running low. Only %d left in stock", stock.StockName, stock.Quantity),
					IsRead:       false,
					NotiType:     "WARNING",
					ItemID:       stock.StockId,
					ItemName:     stock.StockName,
					NotiItemType: "STOCK",
				})

				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

func (s *OrderService) EditPreOrderStatusNotification(c *fiber.Ctx, orderID string) error {
	preOrderDetail, err := s.GetPreOrderOrderDetail(c, orderID)
	if err != nil {
		return err
	}

	stocks, err := s.stockService.GetAllStocks(c, preOrderDetail.UserID)
	if err != nil {
		return err
	}

	recipeMap := make(map[string]string)

	for _, item := range preOrderDetail.OrderProduct {
		recipeMap[item.RecipeID] = item.RecipeID
	}

	for _, stock := range stocks.Stocks {

		if stock.StockId == recipeMap[stock.StockId] && stock.Quantity < stock.StockLessThan {

			deviceToken, err := s.userRepo.GetDeviceToken(preOrderDetail.UserID)
			if err != nil {
				return err
			}

			if stock.Quantity == 0 {
				err = config.SendToToken(
					s.firebaseApp,
					*deviceToken,
					"Restock Reminder!",
					fmt.Sprintf("%s is running out", stock.StockName),
				)
				if err != nil {
					return err
				}

				err = s.notiService.CreateNotification(c, &domain.CreateNotificationItem{
					UserID:       preOrderDetail.UserID,
					EngTitle:     "Restock Reminder!",
					EngMessage:   fmt.Sprintf("%s is running out", stock.StockName),
					IsRead:       false,
					NotiType:     "ALERT",
					ItemID:       stock.StockId,
					ItemName:     stock.StockName,
					NotiItemType: "STOCK",
				})
				if err != nil {
					return err
				}
			} else {
				err = config.SendToToken(
					s.firebaseApp,
					*deviceToken,
					"Stock Up Time!",
					fmt.Sprintf("%s is running low. Only %d left in stock", stock.StockName, stock.Quantity),
				)
				if err != nil {
					return err
				}
				err = s.notiService.CreateNotification(c, &domain.CreateNotificationItem{
					UserID:       preOrderDetail.UserID,
					EngTitle:     "Stock Up Time!",
					EngMessage:   fmt.Sprintf("%s is running low. Only %d left in stock", stock.StockName, stock.Quantity),
					IsRead:       false,
					NotiType:     "WARNING",
					ItemID:       stock.StockId,
					ItemName:     stock.StockName,
					NotiItemType: "STOCK",
				})

				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}

func (s *OrderService) BeforePickUpPreOrderNotifiation() error {

	users, err := s.userService.GetAllUsers()
	if err != nil {
		return err
	}

	for _, userId := range users {
		deviceToken, err := s.userRepo.GetDeviceToken(userId)
		if err != nil {
			return err
		}

		orderList, err := s.GetAllOrders(nil, userId)
		if err == nil {
			for _, order := range orderList.Orders {
				if order.IsPreOrder && order.OrderStatus != "DONE" {
					now := time.Now()
					layout := "02/01/2006 03:04 PM"
					pickUpDate, _ := time.ParseInLocation(layout, order.PickUpDate, time.Local)
					daysUntilPickUp := int(math.Ceil(pickUpDate.Sub(now).Hours() / 24))

					if daysUntilPickUp == 5 {
						err = config.SendToToken(
							s.firebaseApp,
							*deviceToken,
							"Heads Up! New Order Incoming!",
							fmt.Sprintf("Order #%d for pick-up on %s", order.OrderIndex, order.PickUpDate),
						)
						if err != nil {
							return err
						}

						err = s.notiService.CreateNotification(nil, &domain.CreateNotificationItem{
							UserID:       userId,
							EngTitle:     "Heads Up! New Order Incoming!",
							EngMessage:   fmt.Sprintf("Order #%d for pick-up on %s", order.OrderIndex, order.PickUpDate),
							IsRead:       false,
							NotiType:     "INFO",
							ItemID:       order.OrderID,
							ItemName:     order.PickUpDate,
							NotiItemType: "ORDER",
						})
						if err != nil {
							return err
						}
						break
					} else if daysUntilPickUp == 3 {
						err = config.SendToToken(
							s.firebaseApp,
							*deviceToken,
							"Almost Time! Upcoming Order in 3 Days!",
							fmt.Sprintf("Order #%d for pick-up on %s", order.OrderIndex, order.PickUpDate),
						)
						if err != nil {
							return err
						}

						err = s.notiService.CreateNotification(nil, &domain.CreateNotificationItem{
							UserID:       userId,
							EngTitle:     "Almost Time! Upcoming Order in 3 Days!",
							EngMessage:   fmt.Sprintf("Order #%d for pick-up on %s", order.OrderIndex, order.PickUpDate),
							IsRead:       false,
							NotiType:     "WARNING",
							ItemID:       order.OrderID,
							ItemName:     order.PickUpDate,
							NotiItemType: "ORDER",
						})
						if err != nil {
							return err
						}
					} else if daysUntilPickUp == 1 {
						err = config.SendToToken(
							s.firebaseApp,
							*deviceToken,
							"Final Prep! Order Pickup Tomorrow!",
							fmt.Sprintf("Order #%d for pick-up on %s", order.OrderIndex, order.PickUpDate),
						)
						if err != nil {
							return err
						}

						err = s.notiService.CreateNotification(nil, &domain.CreateNotificationItem{
							UserID:       userId,
							EngTitle:     "Final Prep! Order Pickup Tomorrow!",
							EngMessage:   fmt.Sprintf("Order #%d for pick-up on %s", order.OrderIndex, order.PickUpDate),
							IsRead:       false,
							NotiType:     "ALERT",
							ItemID:       order.OrderID,
							ItemName:     order.PickUpDate,
							NotiItemType: "ORDER",
						})
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
