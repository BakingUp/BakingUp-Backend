package repository

import (
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

func (sr *SettingsRepository) GetFixCost(c *fiber.Ctx, userID string) (*db.UsersModel, error) {
	user, err := sr.db.Users.FindFirst(
		db.Users.UserID.Equals(userID),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (sr *SettingsRepository) ChangeFixCost(c *fiber.Ctx, userFixCost *domain.ChangeFixCostSetting) error {
	_, err := sr.db.Users.FindUnique(
		db.Users.UserID.Equals(userFixCost.UserID),
	).Update(
		db.Users.Rent.Set(userFixCost.Rent),
		db.Users.Salaries.Set(userFixCost.Salaries),
		db.Users.Insurance.Set(userFixCost.Insurance),
		db.Users.Subscriptions.Set(userFixCost.Subscriptions),
		db.Users.Advertising.Set(userFixCost.Advertising),
		db.Users.Electricity.Set(userFixCost.Electricity),
		db.Users.Water.Set(userFixCost.Water),
		db.Users.Gas.Set(userFixCost.Gas),
		db.Users.Other.Set(userFixCost.Other),
		db.Users.Note.Set(userFixCost.Note),
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
