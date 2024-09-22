package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (uh *UserHandler) GetUserInfo(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	
	userInfo, err := uh.svc.GetUserInfo(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get the user information", err.Error())
		return nil
	}

	handleSuccess(c, userInfo)
	return nil
}
