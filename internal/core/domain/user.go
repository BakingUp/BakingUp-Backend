package domain

import "time"

type ExpirationDate struct {
	YellowExpirationDate time.Time `json:"yellow_expiration_date"`
	RedExpirationDate    time.Time `json:"red_expiration_date"`
	BlackExpirationDate  time.Time `json:"black_expiration_date"`
}

type RegisterUserRequest struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Tel         string `json:"tel"`
	StoreName   string `json:"store_name"`
	DeviceToken string `json:"device_token"`
}

type DeviceTokenRequest struct {
	UserID      string `json:"user_id"`
	DeviceToken string `json:"device_token"`
}

//เช็กก่อนว่าต้องมี response ตรงนี้ไหม

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
