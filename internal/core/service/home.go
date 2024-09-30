package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type HomeService struct {
	homeRepo    port.HomeRepository
	userService port.UserService
}

func NewHomeService(homeRepo port.HomeRepository, userService port.UserService) *HomeService {
	return &HomeService{
		homeRepo:    homeRepo,
		userService: userService,
	}
}

func (hs *HomeService) GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error) {
	unreadNotiAmount, err := hs.homeRepo.GetUnreadNotification(c, userID)
	if err != nil {
		return nil, err
	}

	return unreadNotiAmount, nil
}
