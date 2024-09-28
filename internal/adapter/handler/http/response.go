package http

import "github.com/gofiber/fiber/v2"

type response struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func newResponse(status int, message string, data any) response {
	return response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func newSuccessMessageResponse(status int, message string) response {
	return response{
		Status:  status,
		Message: message,
	}
}

func newErrorResponse(status int, message string, err string) response {
	return response{
		Status:  status,
		Message: message,
		Error:   err,
	}
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *fiber.Ctx, data any) {
	rsp := newResponse(200, "Success", data)
	ctx.JSON(rsp)
}

func handleError(ctx *fiber.Ctx, status int, message string, err string) {
	rsp := newErrorResponse(status, message, err)
	ctx.JSON(rsp)
}

// handleSuccessMessage send a success response with the specified status code and custom success message
func handleSuccessMessage(ctx *fiber.Ctx, message string) {
	rsp := newSuccessMessageResponse(200, message)
	ctx.JSON(rsp)
}
