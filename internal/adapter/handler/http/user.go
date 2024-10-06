package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

// EditUserInfo godoc
// @Summary      Edit user information
// @Description  Edit user information by using user information request
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        edit_user_info  body  domain.ManageUserRequest  true  "Edit User Info"
// @Success      200  {object}  response  "Successfully edit the user information."
// @Failure      400  {object}  response     "Cannot edit the user information"
// @Router       /user/editUserInfo [put]
func (uh *UserHandler) EditUserInfo(c *fiber.Ctx) error {
	var editUserInfo domain.ManageUserRequest

	if err := c.BodyParser(&editUserInfo); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}

	if editUserInfo.UserID == "" {
		handleError(c, 400, "UserID is required", "")
		return nil
	}

	err := uh.svc.EditUserInfo(c, &editUserInfo)
	if err != nil {
		handleError(c, 400, "Cannot edit the user information", err.Error())
		return nil
	}

	handleSuccessMessage(c, "Successfully edit the user information.")
	return nil
}
