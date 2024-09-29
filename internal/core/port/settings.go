package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type SettingsRepository interface {
	DeleteAccount(c *fiber.Ctx, userID string) error
	GetLanguage(c *fiber.Ctx, userID string) (*db.UsersModel, error)
	ChangeLanguage(c *fiber.Ctx, userLanguage *domain.ChangeUserLanguage) error
	GetFixCost(c *fiber.Ctx, userID string) (*db.UsersModel, error)
	ChangeFixCost(c *fiber.Ctx, userFixCost *domain.ChangeFixCostSetting) error
	GetColorExpired(c *fiber.Ctx, userID string) (*db.UsersModel, error)
	ChangeColorExpired(c *fiber.Ctx, userColorExpired *domain.ChangeExpirationDateSetting) error
}

type SettingsService interface {
	DeleteAccount(c *fiber.Ctx, userID string) error
	GetLanguage(c *fiber.Ctx, userID string) (*domain.UserLanguage, error)
	ChangeLanguage(c *fiber.Ctx, userLanguage *domain.ChangeUserLanguage) error
	GetFixCost(c *fiber.Ctx, userID string) (*domain.FixCostSetting, error)
	ChangeFixCost(c *fiber.Ctx, userFixCost *domain.ChangeFixCostSetting) error
	GetColorExpired(c *fiber.Ctx, userID string) (*domain.ExpirationDateSetting, error)
	ChangeColorExpired(c *fiber.Ctx, userColorExpired *domain.ChangeExpirationDateSetting) error
}
