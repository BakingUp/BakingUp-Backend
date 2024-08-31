package domain

import "time"

type ExpirationDate struct {
	YellowExpirationDate time.Time `json:"yellow_expiration_date"`
	RedExpirationDate    time.Time `json:"red_expiration_date"`
	BlackExpirationDate  time.Time `json:"black_expiration_date"`
}