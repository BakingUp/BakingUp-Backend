package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

func (s *SettingsService) GetLanguage(c *fiber.Ctx, userID string) (*domain.UserLanguage, error) {
	users, err := s.settingsRepo.GetLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	userLanguageResponse := &domain.UserLanguage{}
	if users.Language == "EN" {
		userLanguageResponse.Language = "English"
	} else {
		userLanguageResponse.Language = "Thai"
	}

	return userLanguageResponse, nil
}

func (s *SettingsService) ChangeLanguage(c *fiber.Ctx, userLanguage *domain.ChangeUserLanguage) error {
	err := s.settingsRepo.ChangeLanguage(c, userLanguage)
	if err != nil {
		return err
	}

	return nil
}
