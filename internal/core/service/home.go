package service

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type HomeService struct {
	homeRepo       port.HomeRepository
	userService    port.UserService
	settingService port.SettingsService
	recipeRepo     port.RecipeRepository
	ingredientRepo port.IngredientRepository
}

func NewHomeService(homeRepo port.HomeRepository, userService port.UserService, settingService port.SettingsService, recipeRepo port.RecipeRepository, ingredientRepo port.IngredientRepository) *HomeService {
	return &HomeService{
		homeRepo:       homeRepo,
		userService:    userService,
		settingService: settingService,
		recipeRepo:     recipeRepo,
		ingredientRepo: ingredientRepo,
	}
}

func (hs *HomeService) GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error) {
	unreadNotiAmount, err := hs.homeRepo.GetUnreadNotification(c, userID)
	if err != nil {
		return nil, err
	}

	return unreadNotiAmount, nil
}

func (hs *HomeService) GetTopProducts(c *fiber.Ctx, userID string, chartType string, saleChannels []string, orderTypes []string) (*domain.FilterProductResponse, error) {
	orders, err := hs.homeRepo.GetTopProducts(c, userID, saleChannels, orderTypes)
	if err != nil {
		return nil, err
	}

	language, err := hs.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	// Initialize the map
	product := make(map[string]int)
	var productList []domain.FilterItemResponse
	productSellingPrice := make(map[string]domain.ProductPricing)

	for _, order := range orders {

		for _, orderProductItem := range order.OrderProducts() {
			recipeName := util.GetRecipeName(orderProductItem.Recipe(), language)

			if _, ok := product[recipeName]; ok {
				product[recipeName] += orderProductItem.ProductQuantity
			} else {
				product[recipeName] = orderProductItem.ProductQuantity
				for _, image := range orderProductItem.Recipe().RecipeImages() {
					if image.RecipeID == orderProductItem.RecipeID {
						productList = append(productList, domain.FilterItemResponse{
							Name: recipeName,
							URL:  image.RecipeURL,
						})
						break
					}
				}
			}
			if stock, ok := orderProductItem.Recipe().Stocks(); ok {
				productSellingPrice[recipeName] = domain.ProductPricing{
					SellingPrice: stock.SellingPrice,
					Cost:         stock.Cost,
				}
			}
		}
	}

	fixCost, _ := hs.settingService.GetFixCost(c, userID)
	var response domain.FilterProductResponse
	for i := 0; i < len(productList); i++ {
		name := productList[i].Name
		quantity := product[name]
		if chartType == "Best Selling" || chartType == "Worst Selling" {
			productList[i].Detail = fmt.Sprintf("%d items sold", quantity)
		} else if chartType == "Top Profit Ratio" || chartType == "Top Profit Revenue" || chartType == "Top Profit Margin" {
			profitRevenue := productSellingPrice[name].SellingPrice * float64(quantity)
			profitProducts := (productSellingPrice[name].SellingPrice - productSellingPrice[name].Cost) * float64(quantity)
			profitMargin := ((profitRevenue - profitProducts) / profitRevenue) * 100
			allFixCost := fixCost.Rent + fixCost.Salaries + fixCost.Insurance + fixCost.Subscriptions + fixCost.Advertising + fixCost.Electricity + fixCost.Water + fixCost.Gas + fixCost.Other
			profitRatio := ((profitProducts - allFixCost) / profitRevenue) * 100
			switch chartType {
			case "Top Profit Revenue":
				productList[i].Detail = fmt.Sprintf("%.2f baht", profitRevenue)
			case "Top Profit Margin":
				productList[i].Detail = fmt.Sprintf("Profit Margin: %.2f%%", profitMargin)
			case "Top Profit Ratio":
				productList[i].Detail = fmt.Sprintf("Profit Ratio: %.2f%%", profitRatio)
			}
		}
		// else if chartType == "Selling Quickly" {

		// }

	}

	if chartType != "Worst Selling" {
		sort.SliceStable(productList, func(i, j int) bool {
			return productList[i].Detail > productList[j].Detail
		})
	} else {
		sort.SliceStable(productList, func(i, j int) bool {
			return productList[i].Detail < productList[j].Detail
		})
	}

	response = domain.FilterProductResponse{
		Products: productList,
	}

	return &response, nil
}

func (hs *HomeService) GetWastedProduct(c *fiber.Ctx, userID string, filterType string, unitType string, sortType string) (*domain.FilterProductResponse, error) {
	var productList []domain.FilterItemResponse

	language, err := hs.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	if filterType == "Wasted Ingredients" {
		ingredientList, err := hs.ingredientRepo.GetAllIngredients(c, userID)
		if err != nil {
			return nil, err
		}

		for _, ingredient := range ingredientList {
			name := util.GetIngredientName(&ingredient, language)
			var productAmount float64
			var productItem domain.FilterItemResponse

			if unitType == "Solid" && (ingredient.Unit == "G" || ingredient.Unit == "KG") {
				productItem.Name = name
				for _, ingredientItem := range ingredient.IngredientDetail() {
					if ingredientItem.ExpirationDate.Before(time.Now()) && ingredient.Unit == "KG" {
						productAmount += ingredientItem.IngredientQuantity
					} else if ingredientItem.ExpirationDate.Before(time.Now()) {
						productAmount += ingredientItem.IngredientQuantity / 1000
					}
				}
			} else if unitType == "Liquid" && (ingredient.Unit == "ML" || ingredient.Unit == "L") {
				for _, ingredientItem := range ingredient.IngredientDetail() {

					if ingredientItem.ExpirationDate.Before(time.Now()) && ingredient.Unit == "L" {
						productAmount += ingredientItem.IngredientQuantity
					} else if ingredientItem.ExpirationDate.Before(time.Now()) {
						productAmount += ingredientItem.IngredientQuantity / 1000
					}
				}
			}
			images := ingredient.IngredientImages()
			if len(images) > 0 {
				productItem.URL = images[0].IngredientURL
			}
			if unitType == "Solid" {
				productItem.Detail = fmt.Sprintf("Wasted Item: %.3f kg", productAmount)
			} else {
				productItem.Detail = fmt.Sprintf("Wasted Item: %.3f l", productAmount)
			}
			productList = append(productList, productItem)
		}
	} else {
		recipes, err := hs.homeRepo.GetTopWastedStock(c, userID)
		if err != nil {
			return nil, err
		}

		for _, recipe := range recipes {
			var productAmount int
			var productItem domain.FilterItemResponse
			name := util.GetRecipeName(&recipe, language)
			productItem.Name = name

			if stock, ok := recipe.Stocks(); ok {
				for _, stockItem := range stock.StockDetail() {
					expiredDate := stockItem.CreatedAt.AddDate(0, 0, stock.DayBeforeExpired.Day())
					if expiredDate.Before(time.Now()) {
						productAmount += stockItem.Quantity
					}

				}
			}

			images := recipe.RecipeImages()
			if len(images) > 0 {
				productItem.URL = images[0].RecipeURL
			}
			productItem.Detail = fmt.Sprintf("Wasted Item: %d", productAmount)
			productList = append(productList, productItem)
		}
	}

	if sortType == "Ascending" {
		sort.SliceStable(productList, func(i, j int) bool {
			return productList[i].Detail < productList[j].Detail
		})
	} else if sortType == "Descending" {
		sort.SliceStable(productList, func(i, j int) bool {
			return productList[i].Detail > productList[j].Detail
		})
	}

	topProducts := productList
	if len(productList) > 5 {
		topProducts = productList[:5]
	}
	response := &domain.FilterProductResponse{
		Products: topProducts,
	}

	return response, nil

}

func (hs *HomeService) GetDashboardChartData(c *fiber.Ctx, userID string) (*domain.DashboardChartDataResponse, error) {
	recipes, err := hs.homeRepo.GetDashboardChartData(c, userID)
	if err != nil {
		return nil, err
	}
	allFixCost, err := hs.settingService.GetFixCost(c, userID)
	if err != nil {
		return nil, err
	}

	language, err := hs.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	var monthOrder = map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	fixCost := allFixCost.Rent + allFixCost.Salaries + allFixCost.Insurance + allFixCost.Subscriptions + allFixCost.Advertising + allFixCost.Electricity + allFixCost.Water + allFixCost.Gas + allFixCost.Other
	costRevenueData := make(map[string]domain.CostRevenueChartItem)
	var allProfit float64
	var profitThreshold []domain.ProfitThresholdChartItem
	for _, recipe := range recipes {
		var recipeProfit float64
		if stock, ok := recipe.Stocks(); ok {
			for _, orderProduct := range recipe.OrderProducts() {
				month := orderProduct.Order().OrderDate.Month().String()
				revenue := stock.SellingPrice * float64(orderProduct.ProductQuantity)
				cost := stock.Cost * float64(orderProduct.ProductQuantity)
				profit := (stock.SellingPrice - stock.Cost) * float64(orderProduct.ProductQuantity)

				_, ok1 := costRevenueData[month]
				// _, ok2 := netProfit[month]
				if ok1 {
					costRevenueData[month] = domain.CostRevenueChartItem{
						Month:     month,
						Revenue:   revenue + costRevenueData[month].Revenue,
						Cost:      cost + costRevenueData[month].Cost,
						NetProfit: profit + costRevenueData[month].NetProfit,
					}
				} else {

					costRevenueItem := domain.CostRevenueChartItem{
						Month:     month,
						Revenue:   revenue,
						Cost:      cost,
						NetProfit: profit - fixCost,
					}

					// netProfitItem := domain.NetProfitChartItem{
					// 	Month:  month,
					// 	Profit: profit - fixCost,
					// }
					// netProfit[month] = netProfitItem
					costRevenueData[month] = costRevenueItem
				}
				recipeProfit += profit
			}
			name := util.GetRecipeName(&recipe, language)
			profitThresholdItem := &domain.ProfitThresholdChartItem{
				Name:      name,
				Threshold: recipeProfit,
			}

			profitThreshold = append(profitThreshold, *profitThresholdItem)
		}
		allProfit += recipeProfit
	}

	var costRevenueResponse []domain.CostRevenueChartItem
	var netProfitResponse []domain.NetProfitChartItem

	for i := 0; i < len(profitThreshold); i++ {
		profitThresholdValue := (profitThreshold[i].Threshold / allProfit) * 100
		profitThreshold[i].Threshold = math.Round(profitThresholdValue*100) / 100
	}

	for month, item := range costRevenueData {
		costRevenueResponse = append(costRevenueResponse, item)
		netProfitItem := &domain.NetProfitChartItem{
			Month:  month,
			Profit: item.NetProfit,
		}
		netProfitResponse = append(netProfitResponse, *netProfitItem)
	}

	sort.Slice(costRevenueResponse, func(i, j int) bool {
		return monthOrder[costRevenueResponse[i].Month] < monthOrder[costRevenueResponse[j].Month]
	})

	sort.Slice(netProfitResponse, func(i, j int) bool {
		return monthOrder[netProfitResponse[i].Month] < monthOrder[netProfitResponse[j].Month]
	})

	response := &domain.DashboardChartDataResponse{
		CostRevenue:     costRevenueResponse,
		NetProfit:       netProfitResponse,
		ProfitThreshold: profitThreshold,
	}

	return response, nil
}
