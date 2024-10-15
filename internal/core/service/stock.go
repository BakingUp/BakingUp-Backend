package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"time"
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
		if stockItem, ok := recipe.Stocks(); ok {
			stock.StockId = stockItem.RecipeID
			for _, stockDetail := range stockItem.StockDetail() {
				if stockDetail.RecipeID == stockItem.RecipeID {
					stock.Quantity += stockDetail.Quantity
				}
			}
			stock.LST = stockItem.Lst
			stock.SellingPrice = stockItem.SellingPrice
		} else {
			continue
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

func (s *StockService) GetAllStocksForOrder(c *fiber.Ctx, userID string) (*domain.OrderStockList, error) {
	stocks, err := s.stockRepo.GetAllStocks(c, userID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	var stockOrderPages []domain.StockOrderPage
	recipe := make(map[string]db.RecipesModel)

	for _, recipeItem := range stocks {
		recipe[recipeItem.RecipeID] = recipeItem
	}

	for _, recipe := range stocks {
		var stockOrderPage domain.StockOrderPage

		stockOrderPage.RecipeID = recipe.RecipeID
		stockOrderPage.RecipeName = util.GetRecipeName(&recipe, language)

		if stockItem, ok := recipe.Stocks(); ok {
			var dateTime time.Time
			for _, stockDetail := range stockItem.StockDetail() {
				stockOrderPage.Quantity += stockDetail.Quantity

				if dateTime.IsZero() {
					dateTime = stockDetail.SellByDate
				} else if stockDetail.SellByDate.Before(dateTime) {
					dateTime = stockDetail.SellByDate
				}
			}
			profit := stockItem.SellingPrice - stockItem.Cost
			stockOrderPage.SellByDate = dateTime.Format("02/01/2006")
			stockOrderPage.SellingPrice = stockItem.SellingPrice
			stockOrderPage.Profit = profit
		} else {
			continue
		}

		for _, recipeImage := range recipe.RecipeImages() {
			if recipeImage.RecipeID == recipe.RecipeID {
				stockOrderPage.RecipeURL = recipeImage.RecipeURL
				break
			}
		}

		stockOrderPages = append(stockOrderPages, stockOrderPage)
	}

	orderStockList := &domain.OrderStockList{
		OrderStocks: stockOrderPages,
	}

	return orderStockList, nil
}

func (s *StockService) GetStockDetail(c *fiber.Ctx, recipeID string) (*domain.StockItemDetail, error) {
	stock, err := s.stockRepo.GetStockDetail(c, recipeID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, stock.Recipe().UserID)
	if err != nil {
		return nil, err
	}

	var stockDetails []domain.StockDetail
	totalQuantity := 0

	for _, stockDetail := range stock.StockDetail() {
		totalQuantity += stockDetail.Quantity

		var detail domain.StockDetail
		detail.StockDetailId = stockDetail.StockDetailID
		detail.CreatedAt = stockDetail.CreatedAt
		detail.LSTStatus = util.CalculateLstStatus(stock.Lst, stockDetail.Quantity)
		detail.Quantity = stockDetail.Quantity
		detail.SellByDate = stockDetail.SellByDate.Format("02/01/2006")

		stockDetails = append(stockDetails, detail)
	}

	var stockURLs []string
	for _, image := range stock.Recipe().RecipeImages() {
		stockURLs = append(stockURLs, image.RecipeURL)
	}

	stockItemDetail := &domain.StockItemDetail{
		StockName:     util.GetRecipeName(stock.Recipe(), language),
		StockURL:      stockURLs,
		Quantity:      totalQuantity,
		LST:           stock.Lst,
		SellingPrice:  stock.SellingPrice,
		StockLessThan: stock.StockLessThan,
		StockDetails:  stockDetails,
	}

	return stockItemDetail, nil
}

func (s *StockService) DeleteStock(c *fiber.Ctx, recipeID string) error {
	err := s.stockRepo.DeleteStock(c, recipeID)
	if err != nil {
		return err
	}

	return nil
}

func (s *StockService) DeleteStockBatch(c *fiber.Ctx, stockDetailID string) error {
	err := s.stockRepo.DeleteStockBatch(c, stockDetailID)
	if err != nil {
		return err
	}

	return nil
}

func (s *StockService) GetStockBatch(c *fiber.Ctx, stockDetailID string) (*domain.StockBatch, error) {
	stockDetail, err := s.stockRepo.GetStockBatch(c, stockDetailID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, stockDetail.Stock().Recipe().UserID)
	if err != nil {
		return nil, err
	}

	images := stockDetail.Stock().Recipe().RecipeImages()
	firstRecipeURL := ""
	if len(images) != 0 {
		firstRecipeURL = images[0].RecipeURL
	}
	stockNote, _ := stockDetail.Note()
	stockBatch := &domain.StockBatch{
		StockDetailId: stockDetail.StockDetailID,
		RecipeName:    util.GetRecipeName(stockDetail.Stock().Recipe(), language),
		RecipeURL:     firstRecipeURL,
		Quantity:      stockDetail.Quantity,
		SellByDate:    stockDetail.SellByDate.Format("02/01/2006"),
		Note:          stockNote,
		NoteCreatedAt: stockDetail.CreatedAt.Format("02/01/2006"),
	}

	return stockBatch, nil
}
