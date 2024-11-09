package repository

import (
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

//func (sr *StockRepository) GetAllStocksForOrder(c *fiber.Ctx, userID string) ([]db.RecipesModel, error) {
//	recipes, err := sr.db.Recipes.FindMany(
//		db.Recipes.UserID.Equals(userID),
//	).With(
//		db.Recipes.Stocks.Fetch().With(
//			db.Stocks.StockDetail.Fetch(),
//		),
//		db.Recipes.RecipeImages.Fetch(),
//	).Exec(c.Context())
//
//	if err != nil {
//		return nil, err
//	}
//
//	return recipes, nil
//
//}

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

func (sr *StockRepository) DeleteStockBatch(c *fiber.Ctx, stockDetailID string) error {
	_, err := sr.db.StockDetail.FindMany(
		db.StockDetail.RecipeID.Equals(stockDetailID),
	).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *StockRepository) GetStockBatch(c *fiber.Ctx, stockDetailID string) (*db.StockDetailModel, error) {
	stockDetail, err := sr.db.StockDetail.FindFirst(
		db.StockDetail.StockDetailID.Equals(stockDetailID),
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

func (sr *StockRepository) AddStock(c *fiber.Ctx, stock *domain.AddStockPayload) error {
	_, err := sr.db.Stocks.CreateOne(
		db.Stocks.Recipe.Link(
			db.Recipes.RecipeID.Equals(stock.StockID),
		),
		db.Stocks.Lst.Set(stock.LST),
		db.Stocks.SellingPrice.Set(stock.SellingPrice),
		// TODO: Remove
		db.Stocks.Cost.Set(0),
		db.Stocks.StockLessThan.Set(stock.StockLessThan),
		db.Stocks.DayBeforeExpired.Set(stock.ExpirationDate),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *StockRepository) AddStockDetail(c *fiber.Ctx, stockDetail *domain.AddStockDetailPayload) error {
	_, err := sr.db.StockDetail.CreateOne(
		db.StockDetail.StockDetailID.Set(stockDetail.StockDetailID),
		db.StockDetail.Stock.Link(
			db.Stocks.RecipeID.Equals(stockDetail.RecipeID),
		),
		db.StockDetail.CreatedAt.Set(time.Now()),
		db.StockDetail.SellByDate.Set(stockDetail.SellByDate),
		db.StockDetail.Quantity.Set(int(stockDetail.Quantity)),
		db.StockDetail.Note.Set(stockDetail.Note),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *StockRepository) EditStock(c *fiber.Ctx, stock *domain.EditStockPayload) error {
	_, err := sr.db.Stocks.FindUnique(
		db.Stocks.RecipeID.Equals(stock.RecipeID),
	).Update(
		db.Stocks.Lst.Set(stock.LST),
		db.Stocks.SellingPrice.Set(stock.SellingPrice),
		db.Stocks.StockLessThan.Set(stock.StockLessThan),
		db.Stocks.DayBeforeExpired.Set(stock.ExpirationDate),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (sr *StockRepository) GetEditStockDetail(c *fiber.Ctx, recipeID string) (*db.StocksModel, error) {
	stock, err := sr.db.Stocks.FindFirst(
		db.Stocks.RecipeID.Equals(recipeID),
	).With(
		db.Stocks.Recipe.Fetch().With(
			db.Recipes.RecipeImages.Fetch(),
		),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return stock, nil
}