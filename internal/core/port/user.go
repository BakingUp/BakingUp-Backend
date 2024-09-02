package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetUser(c *fiber.Ctx, userID string) (*db.UsersModel, error)
	CreateUser(user *domain.RegisterUserRequest) error
	AddDeviceToken(req *domain.DeviceTokenRequest) error
}

type UserService interface {
	GetUserLanguage(c *fiber.Ctx, userID string) (*db.Language, error)
	GetUserExpirationDate(c *fiber.Ctx, userID string) (*domain.ExpirationDate, error)
}
