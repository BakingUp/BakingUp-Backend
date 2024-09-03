package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req domain.RegisterUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.UserResponse{
			Status:  400,
			Message: "Invalid request body.",
		})
	}

	response, err := h.userService.RegisterUser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *AuthHandler) AddDeviceToken(c *fiber.Ctx) error {
	var req domain.DeviceTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.UserResponse{
			Status:  400,
			Message: "Invalid request body.",
		})
	}

	response, err := h.userService.AddDeviceToken(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *AuthHandler) DeleteDeviceToken(c *fiber.Ctx) error {
	var req domain.DeviceTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.UserResponse{
			Status:  400,
			Message: "Invalid request body.",
		})
	}

	response, err := h.userService.DeleteDeviceToken(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthHandler) DeleteAllExceptDeviceToken(c *fiber.Ctx) error {
	var req domain.DeviceTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.UserResponse{
			Status:  400,
			Message: "Invalid request body.",
		})
	}

	response, err := h.userService.DeleteAllExceptDeviceToken(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
