package repository

import (
	"context"
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserRepository struct {
	db *db.PrismaClient
}

func NewUserRepository(db *db.PrismaClient) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(user *domain.RegisterUserRequest) error {
	ctx := context.Background()

	blackExpirationDate := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	redExpirationDate := time.Date(2000, time.January, 6, 0, 0, 0, 0, time.UTC)
	yellowExpirationDate := time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC)

	createdUser, err := ur.db.Users.CreateOne(
		db.Users.UserID.Set(user.UserID),
		db.Users.FirstName.Set(user.FirstName),
		db.Users.LastName.Set(user.LastName),
		db.Users.Telephone.Set(user.Tel),
		db.Users.StoreName.Set(user.StoreName),
		db.Users.BlackExpirationDate.Set(blackExpirationDate),
		db.Users.RedExpirationDate.Set(redExpirationDate),
		db.Users.YellowExpirationDate.Set(yellowExpirationDate),
		db.Users.Language.Set(db.LanguageEn),
	).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = ur.db.Devices.CreateOne(
		db.Devices.DeviceToken.Set(user.DeviceToken),
		db.Devices.User.Link(
			db.Users.UserID.Equals(createdUser.UserID),
		),
	).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
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

func (ur *UserRepository) AddDeviceToken(req *domain.DeviceTokenRequest) error {
	ctx := context.Background()

	_, err := ur.db.Devices.CreateOne(
		db.Devices.DeviceToken.Set(req.DeviceToken),
		db.Devices.User.Link(
			db.Users.UserID.Equals(req.UserID),
		),
	).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
