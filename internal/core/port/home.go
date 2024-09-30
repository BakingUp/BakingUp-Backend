package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type HomeRepository interface {
	GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error)
}

type HomeService interface {
	GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error)
}
