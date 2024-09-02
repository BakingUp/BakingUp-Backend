package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type StockRepository interface {
	GetAllStocks(c *fiber.Ctx, userID string) ([]db.RecipesModel, error)
}

type StockService interface {
	GetAllStocks(c *fiber.Ctx, userID string) (*domain.StockList, error)
}
