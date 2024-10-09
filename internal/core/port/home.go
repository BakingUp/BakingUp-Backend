package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type HomeRepository interface {
	GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error)
	GetTopProducts(c *fiber.Ctx, userID string, saleChannels []string, orderTypes []string) ([]db.OrdersModel, error)
	GetTopWastedStock(c *fiber.Ctx, userID string) ([]db.RecipesModel, error)
	GetDashboardChartData(c *fiber.Ctx, userID string) ([]db.RecipesModel, error)
}

type HomeService interface {
	GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error)
	GetTopProducts(c *fiber.Ctx, userID string, chartType string, saleChannels []string, orderTypes []string) (*domain.FilterProductResponse, error)
	GetWastedProduct(c *fiber.Ctx, userID string, chartType string, productType string, sortType string) (*domain.FilterProductResponse, error)
	GetDashboardChartData(c *fiber.Ctx, userID string) (*domain.DashboardChartDataResponse, error)
}
