package repository

import (
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type SettingsRepository struct {
	db *db.PrismaClient
}

func NewSettingsRepository(db *db.PrismaClient) *SettingsRepository {
	return &SettingsRepository{
		db,
	}
}

func (sr *SettingsRepository) DeleteAccount(c *fiber.Ctx, userID string) error {
	_, err := sr.db.Users.FindUnique(
		db.Users.UserID.Equals(userID),
	).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *SettingsRepository) GetLanguage(c *fiber.Ctx, userID string) (*db.UsersModel, error) {
	users, err := sr.db.Users.FindFirst(
		db.Users.UserID.Equals(userID),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (sr *SettingsRepository) ChangeLanguage(c *fiber.Ctx, userLanguage *domain.ChangeUserLanguage) error {
	var language db.Language
	if userLanguage.Language == "English" {
		language = "EN"
	} else {
		language = "TH"
	}
	_, err := sr.db.Users.FindUnique(
		db.Users.UserID.Equals(userLanguage.UserID),
	).Update(
		db.Users.Language.Set(language),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *SettingsRepository) GetFixCost(c *fiber.Ctx, userID string, created_at time.Time) (*db.FixCostModel, error) {
	user, err := sr.db.FixCost.FindFirst(
		db.FixCost.UserID.Equals(userID),
		db.FixCost.CreatedAt.Equals(created_at),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (sr *SettingsRepository) ChangeFixCost(c *fiber.Ctx, userFixCost *domain.ChangeFixCostSetting) error {
	_, err := sr.db.FixCost.FindUnique(
		db.FixCost.FixCostID.Equals(userFixCost.FixCostID),
	).Update(
		db.FixCost.Rent.Set(userFixCost.Rent),
		db.FixCost.Salaries.Set(userFixCost.Salaries),
		db.FixCost.Insurance.Set(userFixCost.Insurance),
		db.FixCost.Subscriptions.Set(userFixCost.Subscriptions),
		db.FixCost.Advertising.Set(userFixCost.Advertising),
		db.FixCost.Electricity.Set(userFixCost.Electricity),
		db.FixCost.Water.Set(userFixCost.Water),
		db.FixCost.Gas.Set(userFixCost.Gas),
		db.FixCost.Other.Set(userFixCost.Other),
		db.FixCost.Note.Set(userFixCost.Note),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *SettingsRepository) GetColorExpired(c *fiber.Ctx, userID string) (*db.UsersModel, error) {
	user, err := sr.db.Users.FindFirst(
		db.Users.UserID.Equals(userID),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (sr *SettingsRepository) ChangeColorExpired(c *fiber.Ctx, userColorExpired *domain.ChangeExpirationDateSetting) error {
	blackExpirationDate := time.Date(2000, 1, userColorExpired.BlackExpirationDate+1, 0, 0, 0, 0, time.UTC)
	redExpirationDate := time.Date(2000, 1, userColorExpired.RedExpirationDate+1, 0, 0, 0, 0, time.UTC)
	yellowExpirationDate := time.Date(2000, 1, userColorExpired.YellowExpirationDate+1, 0, 0, 0, 0, time.UTC)

	_, err := sr.db.Users.FindUnique(
		db.Users.UserID.Equals(userColorExpired.UserID),
	).Update(
		db.Users.BlackExpirationDate.Set(blackExpirationDate),
		db.Users.RedExpirationDate.Set(redExpirationDate),
		db.Users.YellowExpirationDate.Set(yellowExpirationDate),
	).Exec(c.Context())

	if err != nil {
		return nil
	}

	return nil
}
