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
	context := context.Background()
	if c != nil {
		context = c.Context()
	}
	orders, err := or.db.Orders.FindMany(
		db.Orders.UserID.Equals(userID)).With(db.Orders.CuttingStock.Fetch(), db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch()))).Exec(context)
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
		db.Orders.OrderStatus.Set(db.OrderStatusDone),
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
		).OrderBy(db.StockDetail.SellByDate.Order(db.ASC)).Exec(ctx)
		if err != nil {
			return err
		}
		quantity := product.ProductQuantity
		if quantity > 0 {
			for _, stockDetail := range earliestStockDetail {
				if stockDetail.Quantity > 0 {
					if stockDetail.Quantity >= quantity {
						_, err = or.db.StockDetail.FindUnique(
							db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
						).Update(db.StockDetail.Quantity.Set(stockDetail.Quantity - quantity)).Exec(ctx)
						if err != nil {
							return err
						}

						_, err = or.db.CuttingStock.CreateOne(
							db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
							db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
							db.CuttingStock.Quantity.Set(quantity),
							db.CuttingStock.CuttingTime.Set(time.Now()),
						).Exec(ctx)
						break
					}
					quantity -= stockDetail.Quantity
					_, err = or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
					).Update(db.StockDetail.Quantity.Set(0)).Exec(ctx)
					if err != nil {
						return err
					}

					_, err = or.db.CuttingStock.CreateOne(
						db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
						db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
						db.CuttingStock.Quantity.Set(stockDetail.Quantity),
						db.CuttingStock.CuttingTime.Set(time.Now()),
					).Exec(ctx)
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

		switch preOrderOrder.OrderStatus {
		case "DONE":
			//add cutting stock
			earliestStockDetail, err := or.db.StockDetail.FindMany(
				db.StockDetail.RecipeID.Equals(product.RecipeID),
				db.StockDetail.SellByDate.Gte(time.Now()),
			).OrderBy(db.StockDetail.SellByDate.Order(db.ASC)).Exec(ctx)
			if err != nil {
				return err
			}
			quantity := product.ProductQuantity
			if quantity > 0 {
				for _, stockDetail := range earliestStockDetail {
					if stockDetail.Quantity > 0 {
						if stockDetail.Quantity >= quantity {
							_, err = or.db.StockDetail.FindUnique(
								db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
							).Update(db.StockDetail.Quantity.Set(stockDetail.Quantity - quantity)).Exec(ctx)
							if err != nil {
								return err
							}
							_, err = or.db.CuttingStock.CreateOne(
								db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
								db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
								db.CuttingStock.Quantity.Set(quantity),
								db.CuttingStock.CuttingTime.Set(time.Now()),
							).Exec(ctx)

							break
						}
						quantity -= stockDetail.Quantity
						_, err = or.db.StockDetail.FindUnique(
							db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
						).Update(db.StockDetail.Quantity.Set(0)).Exec(ctx)
						if err != nil {
							return err
						}

						_, err = or.db.CuttingStock.CreateOne(
							db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
							db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
							db.CuttingStock.Quantity.Set(stockDetail.Quantity),
							db.CuttingStock.CuttingTime.Set(time.Now()),
						).Exec(ctx)
					}
				}
			}

		case "IN_PROCESS":
			//add production queue
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

	}

	return nil
}

func (or *OrderRepository) EditOrderStatus(c *fiber.Ctx, orderStatue *domain.EditOrderStatusRequest) error {
	ctx := context.Background()

	order, err := or.db.Orders.FindUnique(
		db.Orders.OrderID.Equals(orderStatue.OrderID),
	).With(db.Orders.OrderProducts.Fetch()).Exec(ctx)
	if err != nil {
		return err
	}

	//Pre-order
	if order.IsPreOrder {
		switch order.OrderStatus {
		case db.OrderStatusInProcess:
			if orderStatue.OrderStatus == "DONE" {
				//case 1 : In-process -> Done
				productionList, err := or.db.ProductionQueue.FindMany(
					db.ProductionQueue.OrderID.Equals(order.OrderID),
				).Exec(ctx)
				if err != nil {
					return err
				}

				// add cutting stock
				for _, production := range productionList {
					earliestStockDetail, err := or.db.StockDetail.FindMany(
						db.StockDetail.RecipeID.Equals(production.RecipeID),
						db.StockDetail.SellByDate.Gte(time.Now()),
					).OrderBy(db.StockDetail.SellByDate.Order(db.ASC)).Exec(ctx)
					if err != nil {
						return err
					}
					quantity := production.ProductionQuantity
					if quantity > 0 {
						for _, stockDetail := range earliestStockDetail {
							if stockDetail.Quantity > 0 {
								if stockDetail.Quantity >= quantity {
									_, err = or.db.StockDetail.FindUnique(
										db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
									).Update(db.StockDetail.Quantity.Set(stockDetail.Quantity - quantity)).Exec(ctx)
									if err != nil {
										return err
									}
									_, err = or.db.CuttingStock.CreateOne(
										db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
										db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
										db.CuttingStock.Quantity.Set(quantity),
										db.CuttingStock.CuttingTime.Set(time.Now()),
									).Exec(ctx)

									break
								}
								quantity -= stockDetail.Quantity
								_, err = or.db.StockDetail.FindUnique(
									db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
								).Update(db.StockDetail.Quantity.Set(0)).Exec(ctx)
								if err != nil {
									return err
								}

								_, err = or.db.CuttingStock.CreateOne(
									db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
									db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
									db.CuttingStock.Quantity.Set(stockDetail.Quantity),
									db.CuttingStock.CuttingTime.Set(time.Now()),
								).Exec(ctx)
							}
						}
					}
				}
				// delete production queue
				_, err = or.db.ProductionQueue.FindMany(
					db.ProductionQueue.OrderID.Equals(order.OrderID),
				).Delete().Exec(ctx)
				if err != nil {
					return err
				}

			} else if orderStatue.OrderStatus == "CANCEL" {
				//case 2 : In-process -> Cancel
				// delete production queue
				_, err = or.db.ProductionQueue.FindMany(
					db.ProductionQueue.OrderID.Equals(order.OrderID),
				).Delete().Exec(ctx)
				if err != nil {
					return err
				}
			}
		case db.OrderStatusDone:
			if orderStatue.OrderStatus == "CANCEL" {
				//case 3 : Done -> Cancel
				// undo cutting stock
				cuttingStockList, err := or.db.CuttingStock.FindMany(
					db.CuttingStock.OrderID.Equals(order.OrderID),
				).Exec(ctx)
				if err != nil {
					return err
				}
				for _, cuttingStock := range cuttingStockList {
					stock, err := or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(cuttingStock.StockDetailID),
					).Exec(ctx)
					if err != nil {
						return err
					}
					newQuantity := stock.Quantity + cuttingStock.Quantity
					_, err = or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(cuttingStock.StockDetailID),
					).Update(db.StockDetail.Quantity.Set(newQuantity)).Exec(ctx)
				}
				// delete cutting stock data
				_, err = or.db.CuttingStock.FindMany(
					db.CuttingStock.OrderID.Equals(order.OrderID),
				).Delete().Exec(ctx)
				if err != nil {
					return err
				}
			} else if orderStatue.OrderStatus == "IN_PROCESS" {
				//case 4 : Done -> In-process
				// undo cutting stock
				cuttingStockList, err := or.db.CuttingStock.FindMany(
					db.CuttingStock.OrderID.Equals(order.OrderID),
				).Exec(ctx)
				if err != nil {
					return err
				}
				for _, cuttingStock := range cuttingStockList {
					stock, err := or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(cuttingStock.StockDetailID),
					).Exec(ctx)
					if err != nil {
						return err
					}
					newQuantity := stock.Quantity + cuttingStock.Quantity
					_, err = or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(cuttingStock.StockDetailID),
					).Update(db.StockDetail.Quantity.Set(newQuantity)).Exec(ctx)
				}
				// delete cutting stock data
				_, err = or.db.CuttingStock.FindMany(
					db.CuttingStock.OrderID.Equals(order.OrderID),
				).Delete().Exec(ctx)
				if err != nil {
					return err
				}
				// add production queue
				for _, product := range order.OrderProducts() {
					_, err = or.db.ProductionQueue.CreateOne(
						db.ProductionQueue.User.Link(db.Users.UserID.Equals(order.UserID)),
						db.ProductionQueue.Recipe.Link(db.Recipes.RecipeID.Equals(product.RecipeID)),
						db.ProductionQueue.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
						db.ProductionQueue.ProductionQuantity.Set(product.ProductQuantity),
					).Exec(ctx)
					if err != nil {
						return err
					}
				}
			}
		case db.OrderStatusCancel:
			if orderStatue.OrderStatus == "DONE" {
				//case 5 : Cancel -> Done
				// add cutting stock
				for _, product := range order.OrderProducts() {
					earliestStockDetail, err := or.db.StockDetail.FindMany(
						db.StockDetail.RecipeID.Equals(product.RecipeID),
						db.StockDetail.SellByDate.Gte(time.Now()),
					).OrderBy(db.StockDetail.SellByDate.Order(db.ASC)).Exec(ctx)
					if err != nil {
						return err
					}
					quantity := product.ProductQuantity
					if quantity > 0 {
						for _, stockDetail := range earliestStockDetail {
							if stockDetail.Quantity > 0 {
								if stockDetail.Quantity >= quantity {
									_, err = or.db.StockDetail.FindUnique(
										db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
									).Update(db.StockDetail.Quantity.Set(stockDetail.Quantity - quantity)).Exec(ctx)
									if err != nil {
										return err
									}
									_, err = or.db.CuttingStock.CreateOne(
										db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
										db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
										db.CuttingStock.Quantity.Set(quantity),
										db.CuttingStock.CuttingTime.Set(time.Now()),
									).Exec(ctx)

									break
								}
								quantity -= stockDetail.Quantity
								_, err = or.db.StockDetail.FindUnique(
									db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
								).Update(db.StockDetail.Quantity.Set(0)).Exec(ctx)
								if err != nil {
									return err
								}

								_, err = or.db.CuttingStock.CreateOne(
									db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
									db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
									db.CuttingStock.Quantity.Set(stockDetail.Quantity),
									db.CuttingStock.CuttingTime.Set(time.Now()),
								).Exec(ctx)
							}
						}
					}
				}

			} else if orderStatue.OrderStatus == "IN_PROCESS" {
				//case 6 : Cancel -> In-process
				// add production queue
				for _, product := range order.OrderProducts() {
					_, err = or.db.ProductionQueue.CreateOne(
						db.ProductionQueue.User.Link(db.Users.UserID.Equals(order.UserID)),
						db.ProductionQueue.Recipe.Link(db.Recipes.RecipeID.Equals(product.RecipeID)),
						db.ProductionQueue.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
						db.ProductionQueue.ProductionQuantity.Set(product.ProductQuantity),
					).Exec(ctx)
					if err != nil {
						return err
					}
				}
			}
		}

	} else {
		//in-store
		switch order.OrderStatus {
		case db.OrderStatusDone:
			if orderStatue.OrderStatus == "CANCEL" {
				// case 1 : Done -> Cancel
				// undo cutting stock
				cuttingStockList, err := or.db.CuttingStock.FindMany(
					db.CuttingStock.OrderID.Equals(order.OrderID),
				).Exec(ctx)
				if err != nil {
					return err
				}
				for _, cuttingStock := range cuttingStockList {
					stock, err := or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(cuttingStock.StockDetailID),
					).Exec(ctx)
					if err != nil {
						return err
					}
					newQuantity := stock.Quantity + cuttingStock.Quantity
					_, err = or.db.StockDetail.FindUnique(
						db.StockDetail.StockDetailID.Equals(cuttingStock.StockDetailID),
					).Update(db.StockDetail.Quantity.Set(newQuantity)).Exec(ctx)
				}
				// delete cutting stock data
				_, err = or.db.CuttingStock.FindMany(
					db.CuttingStock.OrderID.Equals(order.OrderID),
				).Delete().Exec(ctx)
				if err != nil {
					return err
				}
			}
		case db.OrderStatusCancel:
			if orderStatue.OrderStatus == "DONE" {
				// case 2 : Cancel -> Done
				// add cutting stock
				for _, product := range order.OrderProducts() {
					earliestStockDetail, err := or.db.StockDetail.FindMany(
						db.StockDetail.RecipeID.Equals(product.RecipeID),
						db.StockDetail.SellByDate.Gte(time.Now()),
					).OrderBy(db.StockDetail.SellByDate.Order(db.ASC)).Exec(ctx)
					if err != nil {
						return err
					}
					quantity := product.ProductQuantity
					if quantity > 0 {
						for _, stockDetail := range earliestStockDetail {
							if stockDetail.Quantity > 0 {
								if stockDetail.Quantity >= quantity {
									_, err = or.db.StockDetail.FindUnique(
										db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
									).Update(db.StockDetail.Quantity.Set(stockDetail.Quantity - quantity)).Exec(ctx)
									if err != nil {
										return err
									}
									_, err = or.db.CuttingStock.CreateOne(
										db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
										db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
										db.CuttingStock.Quantity.Set(quantity),
										db.CuttingStock.CuttingTime.Set(time.Now()),
									).Exec(ctx)

									break
								}
								quantity -= stockDetail.Quantity
								_, err = or.db.StockDetail.FindUnique(
									db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID),
								).Update(db.StockDetail.Quantity.Set(0)).Exec(ctx)
								if err != nil {
									return err
								}

								_, err = or.db.CuttingStock.CreateOne(
									db.CuttingStock.StockDetail.Link(db.StockDetail.StockDetailID.Equals(stockDetail.StockDetailID)),
									db.CuttingStock.Order.Link(db.Orders.OrderID.Equals(order.OrderID)),
									db.CuttingStock.Quantity.Set(stockDetail.Quantity),
									db.CuttingStock.CuttingTime.Set(time.Now()),
								).Exec(ctx)
							}
						}
					}
				}
			}
		}
	}

	_, err = or.db.Orders.FindUnique(
		db.Orders.OrderID.Equals(orderStatue.OrderID),
	).Update(
		db.Orders.OrderStatus.Set(db.OrderStatus(orderStatue.OrderStatus)),
	).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
