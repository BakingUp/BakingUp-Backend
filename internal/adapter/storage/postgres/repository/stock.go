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

func (sr *StockRepository) GetStockDetail(c *fiber.Ctx, recipeID string) (*db.StocksModel, error) {
	stock, err := sr.db.Stocks.FindFirst(
		db.Stocks.RecipeID.Equals(recipeID),
	).With(
		db.Stocks.Recipe.Fetch().With(
			db.Recipes.RecipeImages.Fetch(),
		),
		db.Stocks.StockDetail.Fetch(),
	).Exec((c.Context()))

	if err != nil {
		return nil, err
	}

	return stock, nil
}

func (sr *StockRepository) DeleteStock(c *fiber.Ctx, recipeID string) error {
	_, err := sr.db.Stocks.FindMany(
		db.Stocks.RecipeID.Equals(recipeID),
	).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *StockRepository) GetStockBatch(c *fiber.Ctx, stockDetailID string) (*db.StockDetailModel, error) {
	stockDetail, err := sr.db.StockDetail.FindFirst(
		db.StockDetail.RecipeID.Equals(stockDetailID),
	).With(
		db.StockDetail.Stock.Fetch().With(
			db.Stocks.Recipe.Fetch().With(
				db.Recipes.RecipeImages.Fetch(),
			),
		),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return stockDetail, nil
}
