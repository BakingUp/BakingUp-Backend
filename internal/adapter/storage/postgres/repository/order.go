package repository

import (
	"context"
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type OrderRepository struct {
	db *db.PrismaClient
}

func NewOrderRespository(db *db.PrismaClient) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) GetAllOrders(c *fiber.Ctx, userID string) ([]db.OrdersModel, error) {
	orders, err := or.db.Orders.FindMany(
		db.Orders.UserID.Equals(userID)).With(db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch()))).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (or *OrderRepository) GetOrderDetail(c *fiber.Ctx, orderID string) (*db.OrdersModel, error) {
	order, err := or.db.Orders.FindFirst(
		db.Orders.OrderID.Equals(orderID)).With(db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch()).With(db.Recipes.RecipeImages.Fetch()))).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (or *OrderRepository) DeleteOrder(c *fiber.Ctx, orderID string) error {
	_, err := or.db.Orders.FindUnique(db.Orders.OrderID.Equals(orderID)).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) GetNextOrderIndex(ctx context.Context, userID string) (int, error) {
	maxOrder, err := or.db.Orders.FindFirst(
		db.Orders.UserID.Equals(userID),
	).OrderBy(db.Orders.OrderIndex.Order(db.DESC)).Exec(ctx)

	if err != nil {
		return 0, err
	}

	if maxOrder == nil {
		return 1, nil
	}

	return maxOrder.OrderIndex + 1, nil
}

func (or *OrderRepository) AddInStoreOrder(c *fiber.Ctx, inStoreOrder *domain.AddInStoreOrderRequest) error {
	ctx := context.Background()

	nextOrderIndex, err := or.GetNextOrderIndex(c.Context(), inStoreOrder.UserID)
	if err != nil {
		return err
	}

	timeFormat := "2006-01-02T15:04:05Z07:00"
	createdAt, _ := time.Parse(timeFormat, inStoreOrder.OrderDate)
	noteCreateAt := time.Now()

	if inStoreOrder.NoteText != "" {
		noteCreateAt, _ = time.Parse(timeFormat, inStoreOrder.NoteCreateAt)
	}

	order, err := or.db.Orders.CreateOne(
		db.Orders.User.Link(db.Users.UserID.Equals(inStoreOrder.UserID)),
		db.Orders.OrderPlatform.Set(db.OrderPlatform(inStoreOrder.OrderPlatform)),
		db.Orders.OrderDate.Set(createdAt),
		db.Orders.OrderType.Set(db.OrderType(inStoreOrder.OrderType)),
		db.Orders.IsPreOrder.Set(inStoreOrder.IsPreOrder),
		db.Orders.OrderStatus.Set(db.OrderStatus(inStoreOrder.OrderStatus)),
		db.Orders.OrderIndex.Set(nextOrderIndex),
		db.Orders.OrderNoteText.Set(inStoreOrder.NoteText),
		db.Orders.OrderNoteCreateAt.Set(noteCreateAt),
		db.Orders.OrderTakenBy.Set(inStoreOrder.OrderTakenBy),
	).Exec(ctx)
	if err != nil {
		return err
	}

	for _, product := range inStoreOrder.OrderProducts {
		_, err = or.db.OrderProducts.CreateOne(
			db.OrderProducts.ProductQuantity.Set(product.ProductQuantity),
			db.OrderProducts.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
			db.OrderProducts.Recipe.Link(db.Recipes.RecipeID.Equals(product.RecipeID)),
		).Exec(ctx)
		if err != nil {
			return err
		}

		earliestStockDetail, err := or.db.StockDetail.FindMany(
			db.StockDetail.RecipeID.Equals(product.RecipeID),
			db.StockDetail.SellByDate.Gte(time.Now()),
		).OrderBy(db.StockDetail.CreatedAt.Order(db.ASC)).Exec(ctx)
		if err != nil {
			return err
		}
		quantity := product.ProductQuantity
		if quantity > 0 {
			for _, stockDetail := range earliestStockDetail {
				if stockDetail.Quantity >= quantity {
					_, err = or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
					).Update(db.StockDetail.Quantity.Set(stockDetail.Quantity - quantity)).Exec(ctx)
					if err != nil {
						return err
					}
					break
				}
				quantity -= stockDetail.Quantity
				_, err = or.db.StockDetail.FindUnique(
					db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
				).Update(db.StockDetail.Quantity.Set(0)).Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (or *OrderRepository) AddPreOrderOrder(c *fiber.Ctx, preOrderOrder *domain.AddPreOrderOrderRequest) error {
	ctx := context.Background()

	nextOrderIndex, err := or.GetNextOrderIndex(c.Context(), preOrderOrder.UserID)
	if err != nil {
		return err
	}

	timeFormat := "2006-01-02T15:04:05Z07:00"
	createdAt, _ := time.Parse(timeFormat, preOrderOrder.OrderDate)
	pickUpDateTime, _ := time.Parse(timeFormat, preOrderOrder.PickUpDate)
	noteCreateAt := time.Now()
	customerName := "-"
	phoneNumber := "-"

	if preOrderOrder.NoteText != "" {
		noteCreateAt, _ = time.Parse(timeFormat, preOrderOrder.NoteCreateAt)
	}
	if preOrderOrder.CustomerName != "" {
		customerName = preOrderOrder.CustomerName
	}
	if preOrderOrder.PhoneNumber != "" {
		phoneNumber = preOrderOrder.PhoneNumber
	}

	order, err := or.db.Orders.CreateOne(
		db.Orders.User.Link(db.Users.UserID.Equals(preOrderOrder.UserID)),
		db.Orders.OrderPlatform.Set(db.OrderPlatform(preOrderOrder.OrderPlatform)),
		db.Orders.OrderDate.Set(createdAt),
		db.Orders.OrderType.Set(db.OrderType(preOrderOrder.OrderType)),
		db.Orders.IsPreOrder.Set(preOrderOrder.IsPreOrder),
		db.Orders.OrderStatus.Set(db.OrderStatus(preOrderOrder.OrderStatus)),
		db.Orders.OrderIndex.Set(nextOrderIndex),
		db.Orders.OrderNoteText.Set(preOrderOrder.NoteText),
		db.Orders.OrderNoteCreateAt.Set(noteCreateAt),
		db.Orders.OrderTakenBy.Set(preOrderOrder.OrderTakenBy),
		db.Orders.PickUpDateTime.Set(pickUpDateTime),
		db.Orders.PickUpMethod.Set(db.PickUpMethod(preOrderOrder.PickUpMethod)),
		db.Orders.CustomerName.Set(customerName),
		db.Orders.CustomerPhoneNum.Set(phoneNumber),
	).Exec(ctx)
	if err != nil {
		return err
	}

	for _, product := range preOrderOrder.OrderProducts {
		_, err = or.db.OrderProducts.CreateOne(
			db.OrderProducts.ProductQuantity.Set(product.ProductQuantity),
			db.OrderProducts.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
			db.OrderProducts.Recipe.Link(db.Recipes.RecipeID.Equals(product.RecipeID)),
		).Exec(ctx)
		if err != nil {
			return err
		}

		_, err := or.db.ProductionQueue.CreateOne(
			db.ProductionQueue.User.Link(db.Users.UserID.Equals(preOrderOrder.UserID)),
			db.ProductionQueue.Recipe.Link(db.Recipes.RecipeID.Equals(product.RecipeID)),
			db.ProductionQueue.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
			db.ProductionQueue.ProductionQuantity.Set(product.ProductQuantity),
		).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
