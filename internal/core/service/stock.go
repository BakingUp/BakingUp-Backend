package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type StockService struct {
	stockRepo   port.StockRepository
	userService port.UserService
}

func NewStockService(stockRepo port.StockRepository, userService port.UserService) *StockService {
	return &StockService{
		stockRepo:   stockRepo,
		userService: userService,
	}
}

func (s *StockService) GetAllStocks(c *fiber.Ctx, userID string) (*domain.StockList, error) {
	stocks, err := s.stockRepo.GetAllStocks(c, userID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	var stockItems []domain.StockItem
	recipe := make(map[string]db.RecipesModel)

	for _, recipeItem := range stocks {
		recipe[recipeItem.RecipeID] = recipeItem
	}

	for _, recipe := range stocks {
		var stock domain.StockItem

		stock.StockName = util.GetRecipeName(&recipe, language)

		for _, stockItem := range recipe.Stocks() {
			if stockItem.RecipeID == recipe.RecipeID {

				for _, stockDetail := range stockItem.StockDetail() {
					if stockDetail.RecipeID == stockItem.RecipeID {
						stock.Quantity += stockDetail.Quantity
					}
				}
				stock.LST = stockItem.Lst
				stock.SellingPrice = stockItem.SellingPrice
				break
			}
		}

		for _, recipeImage := range recipe.RecipeImages() {
			if recipeImage.RecipeID == recipe.RecipeID {
				stock.StockURL = recipeImage.RecipeURL
				break
			}
		}

		stock.LSTStatus = util.CalculateLstStatus(stock.LST, stock.Quantity)
		stockItems = append(stockItems, stock)
	}

	stockList := &domain.StockList{
		Stocks: stockItems,
	}

	return stockList, nil

}
