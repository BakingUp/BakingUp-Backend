package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type SettingsService struct {
	settingsRepo port.SettingsRepository
	userService  port.UserService
}

func NewSettingsService(settingsRepo port.SettingsRepository, userService port.UserService) *SettingsService {
	return &SettingsService{
		settingsRepo: settingsRepo,
		userService:  userService,
	}
}

func (s *SettingsService) DeleteAccount(c *fiber.Ctx, userID string) error {
	err := s.settingsRepo.DeleteAccount(c, userID)
	if err != nil {
		return err
	}

	return nil
}
