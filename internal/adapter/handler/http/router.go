package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	router fiber.Router
}

func NewRouter(a *fiber.App, ingredientHandler IngredientHandler, recipeHandler RecipeHandler, authHandler AuthHandler, stockHandler StockHandler, userHandler UserHandler, orderHandler OrderHandler, settingsHandler SettingsHandler, notificationHandler NotificationHandler, homeHandler HomeHandler) (*Router, error) {
	a.Get("/swagger/*", swagger.HandlerDefault)

	api := a.Group("/api")
	{
		user := api.Group("/user")
		{
			user.Get("/getUserInfo", userHandler.GetUserInfo)
			user.Put("/editUserInfo", userHandler.EditUserInfo)
		}
		auth := api.Group("/auth")
		{
			auth.Post("/register", authHandler.Register)
			auth.Post("/addDeviceToken", authHandler.AddDeviceToken)
			auth.Delete("/deleteDeviceToken", authHandler.DeleteDeviceToken)
			auth.Delete("/deleteAllExceptDeviceToken", authHandler.DeleteAllExceptDeviceToken)
		}

		home := api.Group("/home")
		{
			home.Get("/getUnreadNotification", homeHandler.GetUnreadNotification)
			home.Post("/getTopProducts", homeHandler.GetTopProducts)
			home.Get("/getDashboardChartData", homeHandler.GetDashboardChartData)
		}

		ingredient := api.Group("/ingredient")
		{
			ingredient.Get("/getAllIngredients", ingredientHandler.GetAllIngredients)
			ingredient.Get("/getIngredientDetail", ingredientHandler.GetIngredientDetail)
			ingredient.Get("/getIngredientStockDetail", ingredientHandler.GetIngredientStockDetail)
			ingredient.Delete("/deleteIngredientBatchNote", ingredientHandler.DeleteIngredientBatchNote)
			ingredient.Delete("/deleteIngredient", ingredientHandler.DeleteIngredient)
		}

		recipe := api.Group("/recipe")
		{
			recipe.Get("/getAllRecipes", recipeHandler.GetAllRecipes)
			recipe.Get("/getRecipeDetail", recipeHandler.GetRecipeDetail)
			recipe.Delete("/deleteRecipe", recipeHandler.DeleteRecipe)
		}

		stock := api.Group("/stock")
		{
			stock.Get("/getAllStocks", stockHandler.GetAllStocks)
			stock.Get("/getStockDetail", stockHandler.GetStockDetail)
			stock.Delete("/deleteStock", stockHandler.DeleteStock)
			stock.Get("/getStockBatch", stockHandler.GetStockBatch)
			stock.Get("/getAllStocksForOrder", stockHandler.GetAllStocksForOrder)
		}

		order := api.Group("/order")
		{
			order.Get("/getAllOrders", orderHandler.GetAllOrders)
			order.Get("/getOrderDetail", orderHandler.GetOrderDeatil)
			order.Delete("/deleteOrder", orderHandler.DeleteOrder)
			order.Post("/addInStoreOrder", orderHandler.AddInStoreOrder)
		}
		setting := api.Group("/settings")
		{
			setting.Delete("/deleteAccount", settingsHandler.DeleteAccount)
			setting.Get("/getLanguage", settingsHandler.GetLanguage)
			setting.Put("/changeLanguage", settingsHandler.ChangeLanguage)
			setting.Get("/getFixCost", settingsHandler.GetFixCost)
			setting.Put("/changeFixCost", settingsHandler.ChangeFixCost)
			setting.Get("/getColorExpired", settingsHandler.GetColorExpired)
			setting.Put("/changeColorExpired", settingsHandler.ChangeColorExpired)

		}

		notification := api.Group("noti")
		{
			notification.Get("getAllNotifications", notificationHandler.GetAllNotifications)
			notification.Post("createNotification", notificationHandler.CreateNotification)
			notification.Delete("deleteNotification", notificationHandler.DeleteNotification)
			notification.Put("readNotification", notificationHandler.ReadNotification)
			notification.Put("readAllNotifications", notificationHandler.ReadAllNotifications)
		}
	}

	return &Router{api}, nil
}
