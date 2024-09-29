package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetUser(c *fiber.Ctx, userID string) (*db.UsersModel, error)
	CreateUser(user *domain.ManageUserRequest) error
	AddDeviceToken(req *domain.DeviceTokenRequest) error
	DeleteDeviceToken(req *domain.DeviceTokenRequest) error
	DeleteAllExceptDeviceToken(req *domain.DeviceTokenRequest) error
	GetUserProductionQueue(c *fiber.Ctx, userID string) ([]db.OrdersModel, error)
	EditUserInfo(c *fiber.Ctx, editUserRequest *domain.ManageUserRequest) error
}

type UserService interface {
	GetUserLanguage(c *fiber.Ctx, userID string) (*db.Language, error)
	GetUserExpirationDate(c *fiber.Ctx, userID string) (*domain.ExpirationDate, error)
	GetUserInfo(c *fiber.Ctx, userID string) (*domain.UserInfo, error)
	EditUserInfo(c *fiber.Ctx, editUserRequest *domain.ManageUserRequest) error
}
