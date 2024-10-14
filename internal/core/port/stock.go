package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type StockRepository interface {
	GetAllStocks(c *fiber.Ctx, userID string) ([]db.RecipesModel, error)
	GetStockDetail(c *fiber.Ctx, recipeID string) (*db.StocksModel, error)
	DeleteStock(c *fiber.Ctx, recipeID string) error
	DeleteStockBatch(c *fiber.Ctx, stockDetailID string) error
	GetStockBatch(c *fiber.Ctx, stockDetailID string) (*db.StockDetailModel, error)
	//GetAllStocksForOrder(c *fiber.Ctx, userID string) (*db.RecipesModel, error)
}

type StockService interface {
	GetAllStocks(c *fiber.Ctx, userID string) (*domain.StockList, error)
	GetStockDetail(c *fiber.Ctx, recipeID string) (*domain.StockItemDetail, error)
	DeleteStock(c *fiber.Ctx, recipeID string) error
	DeleteStockBatch(c *fiber.Ctx, stockDetailID string) error
	GetStockBatch(c *fiber.Ctx, stockDetailID string) (*domain.StockBatch, error)
	GetAllStocksForOrder(c *fiber.Ctx, userID string) (*domain.OrderStockList, error)
}
