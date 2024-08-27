package repository

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type UserRepository struct {
	db *db.PrismaClient
}

func NewUserRepository(db *db.PrismaClient) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) GetUser(c *fiber.Ctx, userID string) (*db.UsersModel, error) {
	user, err := ur.db.Users.FindFirst(
		db.Users.UserID.Equals(userID),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return user, nil
}