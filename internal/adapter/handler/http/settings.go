package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type SettingsHandler struct {
	svc port.SettingsService
}

func NewSetingsHandler(svc port.SettingsService) *SettingsHandler {
	return &SettingsHandler{
		svc: svc,
	}
}

// DeleteAccount godoc
// @Summary      Delete an account
// @Description  Delete an account by user id
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  response  "Successfully delete an account"
// @Failure      400  {object}  response     "Cannot delete an account"
// @Router       /settings/deleteAccount [delete]
func (sh *SettingsHandler) DeleteAccount(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	err := sh.svc.DeleteAccount(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot delete an account", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully delete an account")
	return nil
}

// GetLanguage godoc
// @Summary      Get the application language
// @Description  Get the application language by user id
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.UserLanguage  "Success"
// @Failure      400  {object}  response     "Cannot get the language"
// @Router       /settings/getLanguage [get]
func (sh *SettingsHandler) GetLanguage(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	userLanguage, err := sh.svc.GetLanguage(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get the language", err.Error())
		return nil
	}

	handleSuccess(c, userLanguage)
	return nil
}
