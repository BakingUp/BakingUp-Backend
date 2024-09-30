package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type NotificationRepository interface {
	GetAllNotifications(c *fiber.Ctx, userID string) ([]db.NotificationsModel, error)
	CreateNotification(c *fiber.Ctx, notificationItem *domain.CreateNotificationItem) error
	DeleteNotification(c *fiber.Ctx, notiID string) error
	ReadNotification(c *fiber.Ctx, notiID string) error
}

type NotificationService interface {
	GetAllNotifications(c *fiber.Ctx, userID string) (*domain.NotificationList, error)
	CreateNotification(c *fiber.Ctx, notificationItem *domain.CreateNotificationItem) error
	DeleteNotification(c *fiber.Ctx, notiID string) error
	ReadNotification(c *fiber.Ctx, notiID string) error
}
