package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

// ChangeLanguage godoc
// @Summary      Change the application language
// @Description  Change the application language by user id
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        change_language  body  domain.ChangeUserLanguage  true  "Change Language"
// @Success      200  {object}  response  "Successfully change the language."
// @Failure      400  {object}  response     "Cannot change the language"
// @Router       /settings/changeLanguage [put]
func (sh *SettingsHandler) ChangeLanguage(c *fiber.Ctx) error {
	var userLanguage domain.ChangeUserLanguage

	if err := c.BodyParser(&userLanguage); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}

	if userLanguage.UserID == "" {
		handleError(c, 400, "UserID is required", "")
		return nil
	}

	err := sh.svc.ChangeLanguage(c, &userLanguage)
	if err != nil {
		handleError(c, 400, "Cannot change the language", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully change the language.")
	return nil
}

// GetFixCost godoc
// @Summary      Get the fix cost
// @Description  Get the fix cost by user id
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.FixCostSetting  "Success"
// @Failure      400  {object}  response     "Cannot get the fix cost"
// @Router       /settings/getFixCost [get]
func (sh *SettingsHandler) GetFixCost(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	userFixCost, err := sh.svc.GetFixCost(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get the fix cost", err.Error())
	}

	handleSuccess(c, userFixCost)
	return nil
}

// ChangeFixCost godoc
// @Summary      Change the fix cost
// @Description  Change the fix cost by user id
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        change_fix_cost  body  domain.ChangeFixCostSetting  true  "Change Fix Cost"
// @Success      200  {object}  response  "Successfully change the fix cost"
// @Failure      400  {object}  response     "Cannot change the fix cost"
// @Router       /settings/changeFixCost [put]
func (sh *SettingsHandler) ChangeFixCost(c *fiber.Ctx) error {
	var userFixCost domain.ChangeFixCostSetting

	if err := c.BodyParser(&userFixCost); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}

	if userFixCost.UserID == "" {
		handleError(c, 400, "UserID is required", "")
		return nil
	}

	err := sh.svc.ChangeFixCost(c, &userFixCost)
	if err != nil {
		handleError(c, 400, "Cannot change the fix cost", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully change the fix cost")
	return nil
}

// GetColorExpired godoc
// @Summary      Get the color of expired icon
// @Description  Get the color of expired icon by user id
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.ExpirationDateSetting  "Success"
// @Failure      400  {object}  response     "Cannot get the color of expiration icon"
// @Router       /settings/getColorExpired [get]
func (sh *SettingsHandler) GetColorExpired(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	ColorExpiredIcons, err := sh.svc.GetColorExpired(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get the color of expiration icon", err.Error())
		return nil
	}

	handleSuccess(c, ColorExpiredIcons)
	return nil
}
