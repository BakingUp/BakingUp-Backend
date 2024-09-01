package domain

type Stock struct {
	StockID          string `json:"stock_id"`
	StockURL         string `json:"stock_url"`
	Price            string `json:"price"`
	Quantity         string `json:"quantity"`
	ExpirationDate   string `json:"expiration_date"`
	ExpirationStatus string `json:"expiration_status"`
}

type Ingredient struct {
	IngredientId     string `json:"ingredient_id"`
	IngredientName   string `json:"ingredient_name"`
	Quantity         string `json:"quantity"`
	Stock            int    `json:"stock"`
	IngredientURL    string `json:"ingredient_url"`
	ExpirationStatus string `json:"expiration_status"`
}

type IngredientList struct {
	Ingredients []Ingredient `json:"ingredients"`
}

type IngredientDetail struct {
	IngredientName     string   `json:"ingredient_name"`
	IngredientQuantity string   `json:"ingredient_quantity"`
	StockAmount        int      `json:"stock_amount"`
	IngredientURL      []string `json:"ingredient_url"`
	IngredientLessThan int      `json:"ingredient_less_than"`
	Stocks             []Stock  `json:"stocks"`
}
