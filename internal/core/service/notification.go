package service

import (
	"fmt"
	"time"

	firebase "firebase.google.com/go"
	"github.com/BakingUp/BakingUp-Backend/internal/adapter/config"
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type NotificationService struct {
	notificationRepo port.NotificationRepository
	userService      port.UserService
	userRepo         port.UserRepository
	stockService     port.StockService
	firebaseApp      *firebase.App
}

func NewNotificationService(notificationRepo port.NotificationRepository, userService port.UserService, userRepo port.UserRepository, stockService port.StockService, firebaseApp *firebase.App) *NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
		userService:      userService,
		userRepo:         userRepo,
		stockService:     stockService,
		firebaseApp:      firebaseApp,
	}
}

func (ns *NotificationService) GetAllNotifications(c *fiber.Ctx, userID string) (*domain.NotificationList, error) {
	notifications, err := ns.notificationRepo.GetAllNotifications(c, userID)
	if err != nil {
		return nil, err
	}

	language, err := ns.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	var notificationItems []domain.NotificationItem
	for _, item := range notifications {

		notificationItem := &domain.NotificationItem{
			NotiID:       item.NotiID,
			Title:        util.GetNotificationTitle(&item, language),
			Message:      util.GetNotificationMessage(&item, language),
			CreatedAt:    item.CreatedAt.Format(time.RFC3339),
			IsRead:       item.IsRead,
			NotiType:     string(item.NotiType),
			ItemID:       item.ItemID,
			ItemName:     item.ItemName,
			NotiItemType: string(item.NotiItemType),
		}
		notificationItems = append(notificationItems, *notificationItem)
	}

	notificationList := &domain.NotificationList{
		Notis: notificationItems,
	}

	if notificationList.Notis == nil {
		notificationList.Notis = []domain.NotificationItem{}
	}

	return notificationList, err
}

func (ns *NotificationService) CreateNotification(c *fiber.Ctx, notificationItem *domain.CreateNotificationItem) error {
	err := ns.notificationRepo.CreateNotification(c, notificationItem)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NotificationService) DeleteNotification(c *fiber.Ctx, notiID string) error {
	err := ns.notificationRepo.DeleteNotification(c, notiID)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NotificationService) ReadNotification(c *fiber.Ctx, notiID string) error {
	err := ns.notificationRepo.ReadNotification(c, notiID)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NotificationService) ReadAllNotifications(c *fiber.Ctx, userID string) error {
	err := ns.notificationRepo.ReadAllNotifications(c, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NotificationService) AddOrderNotification(c *fiber.Ctx, orderProducts []domain.OrderProduct, userId string) error {
	stocks, err := ns.stockService.GetAllStocks(c, userId)
	if err != nil {
		return err
	}

	recipeMap := make(map[string]string)

	for _, item := range orderProducts {
		recipeMap[item.RecipeID] = item.RecipeID
	}

	for _, stock := range stocks.Stocks {

		if stock.StockId == recipeMap[stock.StockId] && stock.Quantity < stock.StockLessThan {

			deviceToken, err := ns.userRepo.GetDeviceToken(c, userId)
			if err != nil {
				return err
			}

			if stock.Quantity == 0 {
				err = config.SendToToken(
					ns.firebaseApp,
					*deviceToken,
					"Restock Reminder!",
					fmt.Sprintf("%s is running out", stock.StockName),
				)
				if err != nil {
					return err
				}

				err = ns.CreateNotification(c, &domain.CreateNotificationItem{
					UserID:       userId,
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
					ns.firebaseApp,
					*deviceToken,
					"Stock Up Time!",
					fmt.Sprintf("%s is running low. Only %d left in stock", stock.StockName, stock.Quantity),
				)
				if err != nil {
					return err
				}
				err = ns.CreateNotification(c, &domain.CreateNotificationItem{
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

func (ns *NotificationService) preOrderNotification(c *fiber.Ctx, orderProducts []domain.OrderProduct, userID string) error {

	return nil
}
