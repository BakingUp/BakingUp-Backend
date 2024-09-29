package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type NotificationHandler struct {
	svc port.NotificationService
}

func NewNotificationHandler(svc port.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		svc: svc,
	}
}

// GetAllNotifications godoc
// @Summary      Get all notifications
// @Description  Get all notifications by user ID
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.NotificationList  "Success"
// @Failure      400  {object}  response     "Cannot get all notifications of the user."
// @Router       /noti/getAllNotifications [get]
func (nh *NotificationHandler) GetAllNotifications(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	notifications, err := nh.svc.GetAllNotifications(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all notifications of the user.", err.Error())
		return nil
	}

	handleSuccess(c, notifications)
	return nil
}
