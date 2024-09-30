package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

// CreateNotification godoc
// @Summary      Create a new notification
// @Description  Create a new notification by user id
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        notification_item  body  domain.CreateNotificationItem  true  "Notification Item"
// @Success      200  {object}  response  "Successfully add a new notification."
// @Failure      400  {object}  response     "Cannot add a new notification."
// @Router       /noti/createNotification [post]
func (nh *NotificationHandler) CreateNotification(c *fiber.Ctx) error {
	var notificationItem domain.CreateNotificationItem

	if err := c.BodyParser(&notificationItem); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}

	if notificationItem.UserID == "" {
		handleError(c, 400, "UserID is required", "")
		return nil
	}

	err := nh.svc.CreateNotification(c, &notificationItem)
	if err != nil {
		handleError(c, 400, "Cannot add a new notification.", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully add a new notification.")
	return nil
}

// DeleteNotification godoc
// @Summary      Delete a notification
// @Description  Delete a notification by notification id
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        noti_id  query  string  true  "Noti ID"
// @Success      200  {object}  response  "Successfully delete a notification."
// @Failure      400  {object}  response     "Cannot delete a notification."
// @Router       /noti/deleteNotification [delete]
func (nh *NotificationHandler) DeleteNotification(c *fiber.Ctx) error {
	notiID := c.Query("noti_id")

	err := nh.svc.DeleteNotification(c, notiID)
	if err != nil {
		handleError(c, 400, "Cannot delete a notification.", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully delete a notification.")
	return nil
}

// ReadNotification godoc
// @Summary      Read a notification message
// @Description  Read a notification message by notification id
// @Tags         notification
// @Accept       json
// @Produce      json
// @Param        noti_id  query  string  true  "Noti ID"
// @Success      200  {object}  response  "Successfully update the read status of the notification."
// @Failure      400  {object}  response     "Cannot update the read status of the notification."
// @Router       /noti/readNotification [put]
func (nh *NotificationHandler) ReadNotification(c *fiber.Ctx) error {
	notiID := c.Query("noti_id")

	err := nh.svc.ReadNotification(c, notiID)
	if err != nil {
		handleError(c, 400, "Cannot update the read status of the notification.", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully update the read status of the notification.")
	return nil
}
