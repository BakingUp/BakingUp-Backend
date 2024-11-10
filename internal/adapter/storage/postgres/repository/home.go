package repository

import (
	"strconv"
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type HomeRepository struct {
	db *db.PrismaClient
}

func NewHomeRepository(db *db.PrismaClient) *HomeRepository {
	return &HomeRepository{
		db,
	}
}

func (hr *HomeRepository) GetUnreadNotification(c *fiber.Ctx, userID string) (*domain.UnreadNotification, error) {
	var unReadNotificationAmount domain.UnreadNotification
	var result []map[string]string
	err := hr.db.Prisma.QueryRaw(`SELECT COUNT(*)  FROM "Notifications" WHERE user_id = $1 AND is_read = $2`, userID, false).Exec(c.Context(), &result)
	if err != nil {
		return nil, err
	}
	count, _ := strconv.Atoi(result[0]["count"])

	unReadNotificationAmount.UnreadNotiAmount = count
	return &unReadNotificationAmount, nil
}

func (hr *HomeRepository) GetTopProducts(c *fiber.Ctx, userID string, saleChannels []string, orderTypes []string) ([]db.OrdersModel, error) {
	var orderTypeValues []db.OrderType
	for _, orderType := range orderTypes {
		orderTypeValues = append(orderTypeValues, db.OrderType(orderType))
	}

	var saleChannelValues []db.OrderPlatform
	for _, saleChannel := range saleChannels {
		saleChannelValues = append(saleChannelValues, db.OrderPlatform(saleChannel))
	}

	orders, err := hr.db.Orders.FindMany(
		db.Orders.UserID.Equals(userID),
		db.Orders.OrderType.In(orderTypeValues),
		db.Orders.OrderPlatform.In(saleChannelValues),
	).With(
		db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch().With(db.Stocks.StockDetail.Fetch()), db.Recipes.RecipeImages.Fetch())),
		db.Orders.User.Fetch().With(db.Users.FixCosts.Fetch()),
		db.Orders.CuttingStock.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (hr *HomeRepository) GetTopWastedStock(c *fiber.Ctx, userID string) ([]db.RecipesModel, error) {
	recipes, err := hr.db.Recipes.FindMany(
		db.Recipes.UserID.Equals(userID),
	).With(
		db.Recipes.OrderProducts.Fetch(),
		db.Recipes.RecipeImages.Fetch(),
		db.Recipes.Stocks.Fetch().With(db.Stocks.StockDetail.Fetch()),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (hr *HomeRepository) GetDashboardChartData(c *fiber.Ctx, userID string, startDateTime time.Time, endDateTime time.Time) ([]db.OrdersModel, error) {
	orders, err := hr.db.Orders.FindMany(
		db.Orders.UserID.Equals(userID),
		db.Orders.OrderDate.AfterEquals(startDateTime),
		db.Orders.OrderDate.BeforeEquals(endDateTime),
	).With(
		db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch())),
	).Exec(c.Context())
	if err != nil {
		return nil, err
	}

	return orders, nil
}
