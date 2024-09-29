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
