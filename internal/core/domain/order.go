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
