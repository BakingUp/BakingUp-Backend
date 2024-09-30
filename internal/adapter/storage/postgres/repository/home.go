package repository

import (
	"strconv"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type HomeRepository struct {
	db *db.PrismaClient
}

func NewHomeRepository(db *db.PrismaClient) *HomeRepository {
	return &HomeRepository{
		db,
	}
}

func (hr *HomeRepository) GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error) {
	var unReadNotificationAmount domain.UnreadNotification
	var result []map[string]string
	err := hr.db.Prisma.QueryRaw(`SELECT COUNT(*)  FROM "Notifications" WHERE user_id = $1 AND is_read = $2`, userID, false).Exec(c.Context(), &result)
	if err != nil {
		return nil, err
	}
	count, _ := strconv.Atoi(result[0]["count"])

	unReadNotificationAmount.UnreadNotiAmount = count
	return &unReadNotificationAmount, nil
}
