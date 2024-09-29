package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
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

func (s *SettingsService) GetFixCost(c *fiber.Ctx, userID string) (*domain.FixCostSetting, error) {
	user, err := s.settingsRepo.GetFixCost(c, userID)
	if err != nil {
		return nil, err
	}

	rent, _ := user.Rent()
	salaries, _ := user.Salaries()
	insurance, _ := user.Insurance()
	subscriptions, _ := user.Subscriptions()
	advertising, _ := user.Advertising()
	electricity, _ := user.Electricity()
	water, _ := user.Water()
	gas, _ := user.Gas()
	other, _ := user.Other()
	note, _ := user.Note()

	userFixCost := &domain.FixCostSetting{
		Rent:          rent,
		Salaries:      salaries,
		Insurance:     insurance,
		Subscriptions: subscriptions,
		Advertising:   advertising,
		Electricity:   electricity,
		Water:         water,
		Gas:           gas,
		Other:         other,
		Note:          note,
	}

	return userFixCost, nil
}

func (s *SettingsService) ChangeFixCost(c *fiber.Ctx, userFixCost *domain.ChangeFixCostSetting) error {
	err := s.settingsRepo.ChangeFixCost(c, userFixCost)
	if err != nil {
		return err
	}
	return nil
}

func (s *SettingsService) GetColorExpired(c *fiber.Ctx, userID string) (*domain.ExpirationDateSetting, error) {
	user, err := s.settingsRepo.GetColorExpired(c, userID)
	if err != nil {
		return nil, err
	}

	userExpirationDate := &domain.ExpirationDateSetting{
		BlackExpirationDate:  util.DaysSince2000(user.BlackExpirationDate),
		RedExpirationDate:    util.DaysSince2000(user.RedExpirationDate),
		YellowExpirationDate: util.DaysSince2000(user.YellowExpirationDate),
	}

	return userExpirationDate, nil
}

func (s *SettingsService) ChangeColorExpired(c *fiber.Ctx, userColorExpired *domain.ChangeExpirationDateSetting) error {
	err := s.settingsRepo.ChangeColorExpired(c, userColorExpired)
	if err != nil {
		return err
	}

	return nil
}
