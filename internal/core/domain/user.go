package domain

import "time"

type ExpirationDate struct {
	YellowExpirationDate time.Time `json:"yellow_expiration_date"`
	RedExpirationDate    time.Time `json:"red_expiration_date"`
	BlackExpirationDate  time.Time `json:"black_expiration_date"`
}

type ManageUserRequest struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel"`
	StoreName string `json:"store_name"`
}

type DeviceTokenRequest struct {
	UserID      string `json:"user_id"`
	DeviceToken string `json:"device_token"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ProductionQueueItem struct {
	OrderIndex int    `json:"order_index"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	PickUpDate string `json:"pick_up_date"`
	ImgURL     string `json:"recipe_url"`
}

type UserInfo struct {
	FirstName       string                `json:"first_name"`
	LastName        string                `json:"last_name"`
	Tel             string                `json:"tel"`
	StoreName       string                `json:"store_name"`
	ProductionQueue []ProductionQueueItem `json:"production_queue"`
}
