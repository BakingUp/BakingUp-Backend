package repository

import (
	"context"
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

func (ur *UserRepository) CreateUser(user *domain.ManageUserRequest) error {
	ctx := context.Background()

	blackExpirationDate := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	redExpirationDate := time.Date(2000, time.January, 6, 0, 0, 0, 0, time.UTC)
	yellowExpirationDate := time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC)

	_, err := ur.db.Users.CreateOne(
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

	return nil
}

func (ur *UserRepository) GetAllUsers() ([]db.UsersModel, error) {
	context := context.Background()
	users, err := ur.db.Users.FindMany().Exec(context)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUser(c *fiber.Ctx, userID string) (*db.UsersModel, error) {
	context := context.Background()
	if c != nil {
		context = c.Context()
	}
	user, err := ur.db.Users.FindFirst(
		db.Users.UserID.Equals(userID),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetDeviceToken(userID string) (*string, error) {
	context := context.Background()
	deviceToken, err := ur.db.Devices.FindFirst(
		db.Devices.UserID.Equals(userID),
	).Exec(context)
	if err != nil {
		return nil, err
	}

	return &deviceToken.DeviceToken, nil
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

func (ur *UserRepository) DeleteDeviceToken(req *domain.DeviceTokenRequest) error {
	ctx := context.Background()
	_, err := ur.db.Devices.FindMany(
		db.Devices.DeviceToken.Equals(req.DeviceToken),
		db.Devices.UserID.Equals(req.UserID)).Delete().Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteAllExceptDeviceToken(req *domain.DeviceTokenRequest) error {
	ctx := context.Background()

	list, err := ur.db.Devices.FindMany(
		db.Devices.UserID.Equals(req.UserID)).Exec(ctx)

	if err != nil {
		return err
	}

	for _, token := range list {
		if token.DeviceToken != req.DeviceToken {
			_, err := ur.db.Devices.FindUnique(
				db.Devices.DeviceToken.Equals(token.DeviceToken)).Delete().Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ur *UserRepository) GetUserProductionQueue(c *fiber.Ctx, userID string) ([]db.OrdersModel, error) {
	now := time.Now()
	orders, err := ur.db.Orders.FindMany(
		db.Orders.UserID.Equals(userID), db.Orders.PickUpDateTime.Gt(now), db.Orders.IsPreOrder.Equals(true),
	).With(
		db.Orders.ProductionQueue.Fetch().With(db.ProductionQueue.Recipe.Fetch().With(db.Recipes.RecipeImages.Fetch())),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (ur *UserRepository) EditUserInfo(c *fiber.Ctx, editUserRequest *domain.ManageUserRequest) error {
	_, err := ur.db.Users.FindUnique(
		db.Users.UserID.Equals(editUserRequest.UserID),
	).Update(
		db.Users.FirstName.Set(editUserRequest.FirstName),
		db.Users.LastName.Set(editUserRequest.LastName),
		db.Users.Telephone.Set(editUserRequest.Tel),
		db.Users.StoreName.Set(editUserRequest.StoreName),
	).Exec(c.Context())
	if err != nil {
		return err
	}

	return nil
}
