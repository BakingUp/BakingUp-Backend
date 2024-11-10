package http

import (
	"time"

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

	startDateTime := time.Date(filterRequest.StartDateTime.Year(), filterRequest.StartDateTime.Month(), 1, 0, 0, 0, 0, filterRequest.StartDateTime.Location())
	endDateTime := time.Date(filterRequest.EndDateTime.Year(), filterRequest.EndDateTime.Month(), 1, 0, 0, 0, 0, filterRequest.EndDateTime.Location())
	if filterRequest.FilterType != "Wasted Ingredients" && filterRequest.FilterType != "Wasted Bakery Stock" && filterRequest.FilterType != "Selling Quickly" {
		filterResponse, err = hh.homeService.GetTopProducts(c, filterRequest.UserID, filterRequest.FilterType, filterRequest.SalesChannel, filterRequest.OrderTypes, startDateTime, endDateTime)
		if err != nil {
			handleError(c, 400, "Cannot get the filter response.", err.Error())
			return nil
		}
	} else if filterRequest.FilterType == "Selling Quickly" {
		filterResponse, err = hh.homeService.GetProductSellingQuickly(c, filterRequest.UserID, filterRequest.SalesChannel, filterRequest.OrderTypes)
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
// @Param        start_date_time query    string  false "Start Date Time"
// @Param        end_date_time   query    string  false "End Date Time"
// @Success      200  {object}  domain.DashboardChartDataResponse  "Success"
// @Failure      400  {object}  response     "Cannot get data for all charts."
// @Router       /home/getDashboardChartData [get]
func (hh *HomeHandler) GetDashboardChartData(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	startDateTimeStr := c.Query("start_date_time")
	endDateTimeStr := c.Query("end_date_time")

	startDateTime, _ := time.Parse(time.RFC3339, startDateTimeStr)
	endDateTime, _ := time.Parse(time.RFC3339, endDateTimeStr)

	response, err := hh.homeService.GetDashboardChartData(c, userID, startDateTime, endDateTime)
	if err != nil {
		handleError(c, 400, "Cannot get data for all charts.", err.Error())
		return nil
	}

	handleSuccess(c, response)
	return nil
}
