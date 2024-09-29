package repository

import (
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
