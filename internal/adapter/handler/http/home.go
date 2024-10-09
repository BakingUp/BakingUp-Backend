package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

// GetTopProducts godoc
// @Summary      Get top products to display in the intelligent dashboard
// @Description  Get top products to display in the intelligent dashboard by user ID
// @Tags         home
// @Accept       json
// @Produce      json
// @Param        filter_request  body  domain.FilterSellingRequest  true  "Filter Request"
// @Success      200  {object}  domain.FilterProductResponse  "Success"
// @Failure      400  {object}  response     "Cannot get the filter response."
// @Router       /home/getTopProducts [post]
func (hh *HomeHandler) GetTopProducts(c *fiber.Ctx) error {
	var filterRequest domain.FilterSellingRequest

	if err := c.BodyParser(&filterRequest); err != nil {
		handleError(c, 400, "Failed to parse request body", err.Error())
		return nil
	}

	if filterRequest.UserID == "" {
		handleError(c, 400, "UserID is required", "")
		return nil
	}

	var filterResponse *domain.FilterProductResponse
	var err error
	if filterRequest.FilterType != "Wasted Ingredients" && filterRequest.FilterType != "Wasted Bakery Stock" {
		filterResponse, err = hh.homeService.GetTopProducts(c, filterRequest.UserID, filterRequest.FilterType, filterRequest.SalesChannel, filterRequest.OrderTypes)
		if err != nil {
			handleError(c, 400, "Cannot get the filter response.", err.Error())
			return nil
		}
	} else {
		filterResponse, err = hh.homeService.GetWastedProduct(c, filterRequest.UserID, filterRequest.FilterType, filterRequest.UnitType, filterRequest.SortType)
		if err != nil {
			handleError(c, 400, "Cannot get the filter response.", err.Error())
			return nil
		}
	}

	handleSuccess(c, filterResponse)
	return nil
}

// GetDashboardChartData godoc
// @Summary      Get data of each chart on dashboard
// @Description  Get data of each chart on dashboard by user ID
// @Tags         home
// @Accept       json
// @Produce      json
// @Param        user_id  query string  true  "User ID"
// @Success      200  {object}  domain.DashboardChartDataResponse  "Success"
// @Failure      400  {object}  response     "Cannot get data for all charts."
// @Router       /home/getDashboardChartData [get]
func (hh *HomeHandler) GetDashboardChartData(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	response, err := hh.homeService.GetDashboardChartData(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get data for all charts.", err.Error())
		return nil
	}

	handleSuccess(c, response)
	return nil
}
