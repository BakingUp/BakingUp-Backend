package service

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
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
	settingRepo    port.SettingsRepository
	recipeRepo     port.RecipeRepository
	ingredientRepo port.IngredientRepository
}

func NewHomeService(homeRepo port.HomeRepository, userService port.UserService, settingService port.SettingsService, settingRepo port.SettingsRepository, recipeRepo port.RecipeRepository, ingredientRepo port.IngredientRepository) *HomeService {
	return &HomeService{
		homeRepo:       homeRepo,
		userService:    userService,
		settingService: settingService,
		settingRepo:    settingRepo,
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

func (hs *HomeService) GetTopProducts(c *fiber.Ctx, userID string, chartType string, saleChannels []string, orderTypes []string, startDateTime time.Time, endDateTime time.Time) (*domain.FilterProductResponse, error) {
	orders, err := hs.homeRepo.GetTopProducts(c, userID, saleChannels, orderTypes)
	if err != nil {
		return nil, err
	}

	language, err := hs.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

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

	fixCostList, _ := hs.settingRepo.GetFixCost(c, userID, startDateTime, endDateTime)
	var allFixCost float64

	for _, item := range fixCostList {
		rent, _ := item.Rent()
		salaries, _ := item.Salaries()
		insurance, _ := item.Insurance()
		subscriptions, _ := item.Subscriptions()
		advertising, _ := item.Advertising()
		electricity, _ := item.Electricity()
		water, _ := item.Water()
		gas, _ := item.Gas()
		other, _ := item.Other()

		allFixCost += rent + salaries + insurance + subscriptions + advertising + electricity + water + gas + other
	}

	allFixCost = allFixCost / float64(len(fixCostList))
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
			profitRatio := ((profitProducts - allFixCost) / profitRevenue) * 100

			if math.IsNaN(profitMargin) {
				profitMargin = 0
			}

			if math.IsInf(profitRatio, 0) {
				profitRatio = 0
			}
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

	re := regexp.MustCompile(`[-+]?\d*\.?\d+`)

	if chartType != "Worst Selling" {
		sort.Slice(productList, func(i, j int) bool {
			// Extract and convert the numeric value from Detail for product i
			detailI := re.FindString(productList[i].Detail)
			valueI, _ := strconv.ParseFloat(detailI, 64)

			// Extract and convert the numeric value from Detail for product j
			detailJ := re.FindString(productList[j].Detail)
			valueJ, _ := strconv.ParseFloat(detailJ, 64)

			return valueI > valueJ
		})
	} else {
		sort.Slice(productList, func(i, j int) bool {
			// Extract and convert the numeric value from Detail for product i
			detailI := re.FindString(productList[i].Detail)
			valueI, _ := strconv.ParseFloat(detailI, 64)

			// Extract and convert the numeric value from Detail for product j
			detailJ := re.FindString(productList[j].Detail)
			valueJ, _ := strconv.ParseFloat(detailJ, 64)

			return valueI < valueJ
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
				productItem.Detail = fmt.Sprintf("Wasted Item: %.3f L", productAmount)
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

	if sortType == "Ascending Order" {
		sort.SliceStable(productList, func(i, j int) bool {
			return productList[i].Detail < productList[j].Detail
		})
	} else if sortType == "Descending Order" {
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

func (hs *HomeService) GetDashboardChartData(c *fiber.Ctx, userID string, startDateTime time.Time, endDateTime time.Time) (*domain.DashboardChartDataResponse, error) {
	orders, err := hs.homeRepo.GetDashboardChartData(c, userID, startDateTime, endDateTime)
	if err != nil {
		return nil, err
	}

	startDateTimeFilter := time.Date(
		startDateTime.Year(),
		startDateTime.Month(),
		1,
		0, 0, 0, 0,
		startDateTime.Location(),
	)

	endDateTimeFilter := time.Date(
		endDateTime.Year(),
		endDateTime.Month(),
		1,
		0, 0, 0, 0,
		endDateTime.Location(),
	)

	allFixCost, err := hs.settingRepo.GetFixCost(c, userID, startDateTimeFilter, endDateTimeFilter)
	if err != nil {
		return nil, err
	}

	fixCostList := make(map[string]float64)

	for _, item := range allFixCost {
		key := fmt.Sprintf("%d/%02d", item.CreatedAt.Year(), int(item.CreatedAt.Month()))
		rent, _ := item.Rent()
		salaries, _ := item.Salaries()
		insurance, _ := item.Insurance()
		subscriptions, _ := item.Subscriptions()
		advertising, _ := item.Advertising()
		electricity, _ := item.Electricity()
		water, _ := item.Water()
		gas, _ := item.Gas()
		other, _ := item.Other()

		fixCost := rent + salaries + insurance + subscriptions + advertising + electricity + water + gas + other
		fixCostList[key] = fixCost
	}

	language, err := hs.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	costRevenueData := make(map[string]domain.CostRevenueChartItem)
	var allProfit float64
	recipeProfitMap := make(map[string]float64)
	var profitThreshold []domain.ProfitThresholdChartItem
	for _, order := range orders {

		month := fmt.Sprintf("%d/%02d", order.OrderDate.Year(), int(order.OrderDate.Month()))
		for _, orderProduct := range order.OrderProducts() {
			if stock, ok := orderProduct.Recipe().Stocks(); ok {
				revenue := stock.SellingPrice * float64(orderProduct.ProductQuantity)
				cost := stock.Cost * float64(orderProduct.ProductQuantity)
				profit := (stock.SellingPrice - stock.Cost) * float64(orderProduct.ProductQuantity)

				_, ok1 := costRevenueData[month]
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
						Cost:      cost + fixCostList[month],
						NetProfit: profit - fixCostList[month],
					}

					costRevenueData[month] = costRevenueItem
				}

				name := util.GetRecipeName(orderProduct.Recipe(), language)
				_, ok2 := recipeProfitMap[name]

				if ok2 {
					recipeProfitMap[name] += profit
				} else {
					recipeProfitMap[name] = profit
				}

				allProfit += profit
			}
		}
	}

	var costRevenueResponse []domain.CostRevenueChartItem

	for key, value := range recipeProfitMap {
		profitThresholdValue := (value / allProfit) * 100
		profitThresholdValue = math.Round(profitThresholdValue*100) / 100
		profitThresholdItem := &domain.ProfitThresholdChartItem{
			Name:      key,
			Threshold: profitThresholdValue,
		}

		profitThreshold = append(profitThreshold, *profitThresholdItem)
	}

	for _, item := range costRevenueData {
		costRevenueResponse = append(costRevenueResponse, item)

	}

	sort.Slice(costRevenueResponse, func(i, j int) bool {
		return costRevenueResponse[i].Month < costRevenueResponse[j].Month
	})

	response := &domain.DashboardChartDataResponse{
		CostRevenue:     costRevenueResponse,
		ProfitThreshold: profitThreshold,
	}

	return response, nil
}
