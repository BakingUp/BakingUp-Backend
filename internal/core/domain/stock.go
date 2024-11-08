package domain

import "time"

type StockItem struct {
	StockId       string  `json:"stock_id"`
	StockName     string  `json:"stock_name"`
	StockURL      string  `json:"stock_url"`
	Quantity      int     `json:"quantity"`
	LST           int     `json:"lst"`
	SellingPrice  float64 `json:"selling_price"`
	LSTStatus     string  `json:"lst_status"`
	StockLessThan int     `json:"stock_less_than"`
}

type StockList struct {
	Stocks []StockItem `json:"stocks"`
}

type StockDetail struct {
	StockDetailId string    `json:"stock_detail_id"`
	CreatedAt     time.Time `json:"created_at"`
	LSTStatus     string    `json:"lst_status"`
	Quantity      int       `json:"quantity"`
	SellByDate    string    `json:"sell_by_date"`
}

type StockItemDetail struct {
	StockName     string        `json:"stock_name"`
	StockURL      []string      `json:"stock_url"`
	Quantity      int           `json:"quantity"`
	LST           int           `json:"lst"`
	SellingPrice  float64       `json:"selling_price"`
	StockLessThan int           `json:"stock_less_than"`
	StockDetails  []StockDetail `json:"stock_details"`
}

type StockBatch struct {
	StockDetailId string `json:"stock_detail_id"`
	RecipeName    string `json:"recipe_name"`
	RecipeURL     string `json:"recipe_url"`
	Quantity      int    `json:"quantity"`
	SellByDate    string `json:"sell_by_date"`
	Note          string `json:"note"`
	NoteCreatedAt string `json:"note_created_at"`
}

type StockOrderPage struct {
	RecipeID     string  `json:"recipe_id"`
	RecipeName   string  `json:"recipe_name"`
	Quantity     int     `json:"quantity"`
	SellByDate   string  `json:"sell_by_date"`
	RecipeURL    string  `json:"recipe_url"`
	SellingPrice float64 `json:"selling_price"`
	Profit       float64 `json:"profit"`
}

type AddStockRequest struct {
	StockID        string `json:"stock_id"`
	LST            string `json:"lst"`
	ExpirationDate string `json:"expiration_date"`
	SellingPrice   string `json:"selling_price"`
	StockLessThan  string `json:"stock_less_than"`
}

type AddStockPayload struct {
	StockID        string    `json:"stock_id"`
	LST            int       `json:"lst"`
	ExpirationDate time.Time `json:"expiration_date"`
	SellingPrice   float64   `json:"selling_price"`
	StockLessThan  int       `json:"stock_less_than"`
}

type OrderStockList struct {
	OrderStocks []StockOrderPage `json:"order_stocks"`
}

type StockRecipeIngredient struct {
	IngredientID       string  `json:"ingredient_id"`
	IngredientName     string  `json:"ingredient_name"`
	IngredientURL      string  `json:"ingredient_url"`
	IngredientQuantity float64 `json:"ingredient_quantity"`
	StockQuantity      float64 `json:"stock_quantity"`
	Unit               string  `json:"unit"`
}

type StockRecipeDetail struct {
	RecipeName  string                  `json:"recipe_name"`
	TotalTime   string                  `json:"total_time"`
	Servings    int                     `json:"servings"`
	Ingredients []StockRecipeIngredient `json:"ingredients"`
}
