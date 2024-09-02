package repository

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type StockRepository struct {
	db *db.PrismaClient
}

func NewStockRepository(db *db.PrismaClient) *StockRepository {
	return &StockRepository{
		db,
	}
}

func (sr *StockRepository) GetAllStocks(c *fiber.Ctx, userID string) ([]db.RecipesModel, error) {
	recipes, err := sr.db.Recipes.FindMany(
		db.Recipes.UserID.Equals(userID),
	).With(
		db.Recipes.Stocks.Fetch().With(
			db.Stocks.StockDetail.Fetch(),
		),
		db.Recipes.RecipeImages.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return recipes, nil
}
