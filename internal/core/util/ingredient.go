package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BakingUp/BakingUp-Backend/prisma/db"
)

func GetIngredientName(ingredient *db.IngredientsModel, language *db.Language) string {
    if *language == db.LanguageTh {
        return ingredient.IngredientThaiName
    }
	
    return ingredient.IngredientEngName
}

// TODO: daysUntilExpiration should be set according to user's setting
func CalculateExpirationStatus(expirationDate time.Time, blackExpirationDate time.Time, redExpirationDate time.Time, yellowExpirationDate time.Time) string {
	daysUntilExpiration := time.Until(expirationDate).Hours() / 24
    numOfDaysBlack := daysSince2000(blackExpirationDate)
    numOfDaysRed := daysSince2000(redExpirationDate)
    numOfDaysYellow := daysSince2000(yellowExpirationDate)

	switch {
	case daysUntilExpiration <= float64(numOfDaysBlack):
		return "black"
	case daysUntilExpiration <= float64(numOfDaysRed):
		return "red"
	case daysUntilExpiration <= float64(numOfDaysYellow):
		return "yellow"
	case daysUntilExpiration > float64(numOfDaysYellow):
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

func daysSince2000(date time.Time) int {
    startDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
    duration := date.Sub(startDate)
    return int(duration.Hours() / 24)
}
