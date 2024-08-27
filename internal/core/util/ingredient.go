package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BakingUp/BakingUp-Backend/prisma/db"
)

// TODO: daysUntilExpiration should be set according to user's setting
func CalculateExpirationStatus(expirationDate time.Time) string {
	daysUntilExpiration := time.Until(expirationDate).Hours() / 24
	switch {
	case daysUntilExpiration <= 0:
		return "black"
	case daysUntilExpiration <= 5:
		return "red"
	case daysUntilExpiration <= 10:
		return "yellow"
	case daysUntilExpiration >= 10:
		return "green"
	default:
		return "none"
	}
}

func CombineIngredientQuantity(quantity float64, unit db.Unit) string {
    unitStr := ""
    switch unit {
    case db.UnitKg:
        unitStr = "kg"
    case db.UnitG:
        unitStr = "g"
    case db.UnitL:
        unitStr = "l"
    case db.UnitMl:
        unitStr = "ml"
    }

    quantityStr := strconv.FormatFloat(quantity, 'f', -1, 64)
    return fmt.Sprintf("%s %s", quantityStr, unitStr)
}

func CombinePrice(price float64, unit db.Unit) string {
    unitStr := ""
    switch unit {
    case db.UnitKg:
        unitStr = "kg"
    case db.UnitG:
        unitStr = "g"
    case db.UnitL:
        unitStr = "l"
    case db.UnitMl:
        unitStr = "ml"
    }

    priceStr := strconv.FormatFloat(price, 'f', -1, 64)
    return fmt.Sprintf("%s/%s", priceStr, unitStr)
}