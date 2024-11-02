package service

import (
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type NotificationService struct {
	notificationRepo port.NotificationRepository
	userService      port.UserService
}

func NewNotificationService(notificationRepo port.NotificationRepository, userService port.UserService) *NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
		userService:      userService,
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
