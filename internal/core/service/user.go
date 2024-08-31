package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserLanguage(c *fiber.Ctx, userID string) (*db.Language, error) {
	user, err := s.userRepo.GetUser(c, userID)

	if err != nil {
		return nil, err
	}

	return &user.Language, nil
}

func (s *UserService) GetUserExpirationDate(c *fiber.Ctx, userID string) (*domain.ExpirationDate, error) {
	user, err := s.userRepo.GetUser(c, userID)

	if err != nil {
		return nil, err
	}

	expirationDate := &domain.ExpirationDate{
		YellowExpirationDate: user.YellowExpirationDate,
		RedExpirationDate: user.RedExpirationDate,
		BlackExpirationDate: user.BlackExpirationDate,
	}

	return expirationDate, nil
}