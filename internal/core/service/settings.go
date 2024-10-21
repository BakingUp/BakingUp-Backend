package service

import (
	"time"

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

func (s *SettingsService) GetFixCost(c *fiber.Ctx, userID string, created_at time.Time) (*domain.FixCostSetting, error) {
	user, err := s.settingsRepo.GetFixCost(c, userID, created_at, created_at)
	if err != nil {
		return nil, err
	}

	rent, _ := user[0].Rent()
	salaries, _ := user[0].Salaries()
	insurance, _ := user[0].Insurance()
	subscriptions, _ := user[0].Subscriptions()
	advertising, _ := user[0].Advertising()
	electricity, _ := user[0].Electricity()
	water, _ := user[0].Water()
	gas, _ := user[0].Gas()
	other, _ := user[0].Other()
	note, _ := user[0].Note()

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
