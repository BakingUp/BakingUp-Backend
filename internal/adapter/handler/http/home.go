package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	homeService *service.HomeService
}

func NewHomeHandler(homeService *service.HomeService) *HomeHandler {
	return &HomeHandler{
		homeService: homeService,
	}
}

// GetUnreadNotification godoc
// @Summary      Get unread notification amount of user
// @Description  Get unread notification amount of user by user ID
// @Tags         home
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.UnreadNotification  "Success"
// @Failure      400  {object}  response     "Cannot get unread notification amount."
// @Router       /home/getUnreadNotification [get]
func (hh *HomeHandler) GetUnreadNotification(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	unreadNotificationAmount, err := hh.homeService.GetUnreadNotification(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get unread notification amount.", err.Error())
		return nil
	}

	handleSuccess(c, unreadNotificationAmount)
	return nil
}
