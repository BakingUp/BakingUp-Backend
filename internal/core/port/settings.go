package port

import "github.com/gofiber/fiber/v2"

type SettingsRepository interface {
	DeleteAccount(c *fiber.Ctx, userID string) error
}

type SettingsService interface {
	DeleteAccount(c *fiber.Ctx, userID string) error
}
