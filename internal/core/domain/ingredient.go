package domain

import (
	"time"
)

type Stock struct {
	StockID          string `json:"stock_id"`
	StockURL         string `json:"stock_url"`
	Price            string `json:"price"`
	Quantity         string `json:"quantity"`
	ExpirationDate   string `json:"expiration_date"`
	ExpirationStatus string `json:"expiration_status"`
}

type Ingredient struct {
	IngredientId      string    `json:"ingredient_id"`
	IngredientName    string    `json:"ingredient_name"`
	Quantity          string    `json:"quantity"`
	Stock             int       `json:"stock"`
	IngredienLessThan float64   `json:"ingredient_less_than"`
	IngredientURL     string    `json:"ingredient_url"`
	DayBeforeExpire   time.Time `json:"day_before_expire"`
	ExpirationStatus  string    `json:"expiration_status"`
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

type IngredientNote struct {
	IngredientNoteID string `json:"ingredient_note_id"`
	IngredientNote   string `json:"ingredient_note"`
	NoteCreatedAt    string `json:"note_created_at"`
}

type IngredientStockDetail struct {
	IngredientEngName  string           `json:"ingredient_eng_name"`
	IngredientThaiName string           `json:"ingredient_thai_name"`
	IngredientQuantity string           `json:"ingredient_quantity"`
	IngredientPrice    string           `json:"ingredient_price"`
	IngredientBrand    string           `json:"ingredient_brand"`
	IngredientSupplier string           `json:"ingredient_supplier"`
	IngredientStockURL string           `json:"ingredient_stock_url"`
	DayBeforeExpire    string           `json:"day_before_expire"`
	Notes              []IngredientNote `json:"notes"`
}

type AddEditIngredientStockDetail struct {
	IngredientEngName  string `json:"ingredient_eng_name"`
	IngredientThaiName string `json:"ingredient_thai_name"`
	Unit               string `json:"unit"`
}

type AddIngredientRequest struct {
	UserID             string   `json:"user_id"`
	IngredientEngName  string   `json:"ingredient_eng_name"`
	IngredientThaiName string   `json:"ingredient_thai_name"`
	Img                []string `json:"img"`
	Unit               string   `json:"unit"`
	StockLessThan      string   `json:"stock_less_than"`
	DayBeforeExpire    string   `json:"day_before_expire"`
}

type AddIngredientPayload struct {
	UserID             string    `json:"user_id"`
	IngredientID       string    `json:"ingredient_id"`
	IngredientEngName  string    `json:"ingredient_name"`
	IngredientThaiName string    `json:"ingredient_thai_name"`
	StockLessThan      int       `json:"stock_less_than"`
	Unit               string    `json:"unit"`
	DayBeforeExpire    time.Time `json:"day_before_expire"`
}

type AddIngredientImagePayload struct {
	IngredientID string `json:"ingredient_id"`
	ImgUrl       string `json:"img"`
	ImageIndex   string `json:"image_index"`
}

type AddIngredientStockRequest struct {
	UserID          string `json:"user_id"`
	IngredientID    string `json:"ingredient_id"`
	Price           string `json:"price"`
	Quantity        string `json:"quantity"`
	ExpirationDate  string `json:"expiration_date"`
	Supplier        string `json:"supplier"`
	IngredientBrand string `json:"ingredient_brand"`
	Img             string `json:"img"`
	Note            string `json:"note"`
}

type AddIngredientStockPayload struct {
	IngredientStockID  string    `json:"ingredient_stock_id"`
	IngredientID       string    `json:"ingredient_id"`
	IngredientQuantity float64   `json:"ingredient_quantity"`
	Price              float64   `json:"price"`
	ExpirationDate     time.Time `json:"expiration_date"`
	IngredientSupplier string    `json:"ingredient_supplier"`
	IngredientBrand    string    `json:"ingredient_brand"`
	IngredientStockURL string    `json:"ingredient_stock_url"`
	Note               string    `json:"note"`
}

type AddIngredientNotePayload struct {
	IngredientNoteID  string    `json:"ingredient_note_id"`
	IngredientStockID string    `json:"ingredient_stock_id"`
	Note              string    `json:"note"`
	NoteCreatedAt     time.Time `json:"note_created_at"`
}

type EditIngredientRequest struct {
	UserID             string `json:"user_id"`
	IngredientID       string `json:"ingredient_id"`
	IngredientEngName  string `json:"ingredient_eng_name"`
	IngredientThaiName string `json:"ingredient_thai_name"`
	Unit               string `json:"unit"`
	StockLessThan      string `json:"stock_less_than"`
	DayBeforeExpire    string `json:"day_before_expire"`
}

type EditIngredientPayload struct {
	IngredientID       string    `json:"ingredient_id"`
	IngredientEngName  string    `json:"ingredient_name"`
	IngredientThaiName string    `json:"ingredient_thai_name"`
	StockLessThan      int       `json:"stock_less_than"`
	DayBeforeExpire    time.Time `json:"day_before_expire"`
}

type GetAddEditIngredientDetail struct {
	IngredientEngName  string `json:"ingredient_eng_name"`
	IngredientThaiName string `json:"ingredient_thai_name"`
	Unit               string `json:"unit"`
	StockLessThan      string `json:"stock_less_than"`
	DayBeforeExpire    string `json:"day_before_expire"`
}

type GetEditIngredientStockDetail struct {
	IngredientStockID string `json:"ingredient_stock_id"`
	Brand             string `json:"brand"`
	Quantity          string `json:"quantity"`
	Price             string `json:"price"`
	Supplier          string `json:"supplier"`
	ExpirationDate    string `json:"expiration_date"`
}

type EditIngredientStockRequest struct {
	IngredientStockID string `json:"ingredient_stock_id"`
	Brand             string `json:"brand"`
	Quantity          string `json:"quantity"`
	Price             string `json:"price"`
	Supplier          string `json:"supplier"`
	ExpirationDate    string `json:"expiration_date"`
	Note              string `json:"note"`
}

type EditIngredientStockPayload struct {
	IngredientStockID string    `json:"ingredient_stock_id"`
	Brand             string    `json:"brand"`
	Quantity          float64   `json:"quantity"`
	Price             float64   `json:"price"`
	Supplier          string    `json:"supplier"`
	ExpirationDate    time.Time `json:"expiration_date"`
}

type IngredientFromReceiptModel struct {
	IngredientName string `json:"ingredientName"`
	Quantity       string `json:"quantity"`
	Price          string `json:"price"`
}

type IngredientFromReceiptResponse struct {
	IngredientName string `json:"ingredient_name"`
	Quantity       string `json:"quantity"`
	Price          string `json:"price"`
}

type IngredientListFromReceiptModel struct {
	Ingredients []IngredientFromReceiptModel `json:"ingredients"`
}

type IngredientListFromReceiptResponse struct {
	Ingredients []IngredientFromReceiptResponse `json:"ingredients"`
}
