package domain

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"time"
)

type OrderDetail struct {
	OrderID     string         `json:"order_id"`
	OrderIndex  int            `json:"order_index"`
	Total       float64        `json:"total"`
	OrderDate   time.Time      `json:"order_date"`
	PickUpDate  time.Time      `json:"pick_up_date"`
	OrderStatus db.OrderStatus `json:"order_status"`
}

type Orders struct {
	Orders []OrderDetail `json:"orders"`
}
