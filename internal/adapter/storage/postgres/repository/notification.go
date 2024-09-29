package repository

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type NotificationRepository struct {
	db *db.PrismaClient
}

func NewNotificationRepository(db *db.PrismaClient) *NotificationRepository {
	return &NotificationRepository{
		db,
	}
}

func (nr *NotificationRepository) GetAllNotifications(c *fiber.Ctx, userID string) ([]db.NotificationsModel, error) {
	notifications, err := nr.db.Notifications.FindMany(
		db.Notifications.UserID.Equals(userID),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return notifications, nil
}
