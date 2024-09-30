package repository

import (
	"context"
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (nr *NotificationRepository) CreateNotification(c *fiber.Ctx, notificationItem *domain.CreateNotificationItem) error {
	ctx := context.Background()

	timeFormat := "2006-01-02T15:04:05Z07:00"
	createdAt, _ := time.Parse(timeFormat, notificationItem.CreatedAt)

	_, err := nr.db.Notifications.CreateOne(
		db.Notifications.NotiID.Set(uuid.NewString()),
		db.Notifications.User.Link(
			db.Users.UserID.Equals(notificationItem.UserID),
		),
		db.Notifications.EngTitle.Set(notificationItem.EngTitle),
		db.Notifications.ThaiTitle.Set(notificationItem.ThaiTitle),
		db.Notifications.EngMessage.Set(notificationItem.EngMessage),
		db.Notifications.ThaiMessage.Set(notificationItem.ThaiMessage),
		db.Notifications.CreatedAt.Set(createdAt),
		db.Notifications.IsRead.Set(notificationItem.IsRead),
		db.Notifications.NotiType.Set(db.NotificationType(notificationItem.NotiType)),
	).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (nr *NotificationRepository) DeleteNotification(c *fiber.Ctx, notiID string) error {
	_, err := nr.db.Notifications.FindUnique(
		db.Notifications.NotiID.Equals(notiID),
	).Delete().Exec(c.Context())
	if err != nil {
		return err
	}

	return nil
}

func (nr *NotificationRepository) ReadNotification(c *fiber.Ctx, notiID string) error {
	_, err := nr.db.Notifications.FindUnique(
		db.Notifications.NotiID.Equals(notiID),
	).Update(
		db.Notifications.IsRead.Set(true),
	).Exec(c.Context())
	if err != nil {
		return err
	}

	return nil
}
