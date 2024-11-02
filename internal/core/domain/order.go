package domain

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
)

type OrderInfo struct {
	OrderID       string           `json:"order_id"`
	OrderIndex    int              `json:"order_index"`
	OrderPlatform db.OrderPlatform `json:"order_platform"`
	IsPreOrder    bool             `json:"is_pre_order"`
	Total         float64          `json:"total"`
	OrderDate     string           `json:"order_date"`
	PickUpDate    string           `json:"pick_up_date"`
	OrderStatus   db.OrderStatus   `json:"order_status"`
}

type Orders struct {
	Orders []OrderInfo `json:"orders"`
}

type OrderStock struct {
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	StockPrice float64 `json:"stock_price"`
	ImgURL     string  `json:"recipe_url"`
}

type InStoreOrderDetails struct {
	OrderIndex        int              `json:"order_index"`
	OrderStatus       db.OrderStatus   `json:"order_status"`
	OrderPlatform     db.OrderPlatform `json:"order_platform"`
	OrderDate         string           `json:"order_date"`
	OrderTime         string           `json:"order_time"`
	OrderType         db.OrderType     `json:"order_type"`
	OrderTakenBy      string           `json:"order_taken_by"`
	OrderStock        []OrderStock     `json:"order_stock"`
	ToTal             float64          `json:"total"`
	Profit            float64          `json:"profit"`
	OrderNoteText     string           `json:"order_note_text"`
	OrderNoteCreateAt string           `json:"order_note_create_at"`
}

type PreOrderOrderDetails struct {
	OrderIndex        int              `json:"order_index"`
	CustomerName      string           `json:"customer_name"`
	PhoneNumber       string           `json:"phone_number"`
	OrderStatus       db.OrderStatus   `json:"order_status"`
	OrderPlatform     db.OrderPlatform `json:"order_platform"`
	OrderDate         string           `json:"order_date"`
	OrderTime         string           `json:"order_time"`
	OrderType         db.OrderType     `json:"order_type"`
	PickUpDate        string           `json:"pick_up_date"`
	PickUpTime        string           `json:"pick_up_time"`
	PickUpMethod      db.PickUpMethod  `json:"pick_up_method"`
	OrderTakenBy      string           `json:"order_taken_by"`
	OrderStock        []OrderStock     `json:"order_stock"`
	ToTal             float64          `json:"total"`
	Profit            float64          `json:"profit"`
	OrderNoteText     string           `json:"order_note_text"`
	OrderNoteCreateAt string           `json:"order_note_create_at"`
}

type OrderProduct struct {
	RecipeID        string `json:"recipe_id"`
	ProductQuantity int    `json:"product_quantity"`
}

type AddInStoreOrderRequest struct {
	UserID        string         `json:"user_id"`
	OrderStatus   string         `json:"order_status"`
	OrderPlatform string         `json:"order_platform"`
	OrderDate     string         `json:"order_date"`
	OrderType     string         `json:"order_type"`
	OrderTakenBy  string         `json:"order_taken_by"`
	OrderProducts []OrderProduct `json:"order_stocks"`
	IsPreOrder    bool           `json:"is_pre_order"`
	NoteText      string         `json:"note_text,omitempty"`
	NoteCreateAt  string         `json:"note_create_at,omitempty"`
}

type AddPreOrderOrderRequest struct {
	UserID        string         `json:"user_id"`
	OrderStatus   string         `json:"order_status"`
	OrderPlatform string         `json:"order_platform"`
	OrderDate     string         `json:"order_date"`
	OrderType     string         `json:"order_type"`
	OrderTakenBy  string         `json:"order_taken_by"`
	CustomerName  string         `json:"customer_name,omitempty"`
	PhoneNumber   string         `json:"phone_number,omitempty"`
	PickUpDate    string         `json:"pick_up_date"`
	PickUpMethod  string         `json:"pick_up_method"`
	OrderProducts []OrderProduct `json:"order_stocks"`
	IsPreOrder    bool           `json:"is_pre_order"`
	NoteText      string         `json:"note_text,omitempty"`
	NoteCreateAt  string         `json:"note_create_at,omitempty"`
}

type EditOrderStatusRequest struct {
	OrderID     string `json:"order_id"`
	OrderStatus string `json:"order_status"`
}
