package util

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	numOfDaysBlack := DaysSince2000(blackExpirationDate)
	numOfDaysRed := DaysSince2000(redExpirationDate)
	numOfDaysYellow := DaysSince2000(yellowExpirationDate)

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

func CombinePrice(price float64, unit db.Unit, quantity float64) string {
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

    actualPrice := price / quantity
    priceStr := strconv.FormatFloat(actualPrice, 'f', 4, 64)
    priceStr = strings.TrimRight(strings.TrimRight(priceStr, "0"), ".")
    return fmt.Sprintf("%s/%s", priceStr, unitStr)
}

func DaysSince2000(date time.Time) int {
	startDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	duration := date.Sub(startDate)
	return int(duration.Hours() / 24)
}

func ExpirationDate(daysSince2000 string) time.Time {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	days, _ := strconv.Atoi(daysSince2000)
	expirationDate := baseDate.AddDate(0, 0, days)
	return expirationDate
}

func UploadIngredientImage(userId string, ingredientId string, imgBase64 string, index string) (string, error) {
	// Decode the base64 string to a byte slice
	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", err
	}
	// Create the path based on userId, and ingredientId
	filePath := filepath.Join(fmt.Sprintf("images/%s/ingredients/%s/%s.jpg", userId, ingredientId, index))

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return "", err
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write the byte slice to the file
	if _, err := file.Write(imgBytes); err != nil {
		return "", err
	}

	// Return the file path as the image URL
	return filePath, nil
}

func UploadIngredientStockImage(userId string, ingredientId string, ingredientStockId string, imgBase64 string) (string, error) {
	// Decode the base64 string to a byte slice
	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", err
	}

	// Create the path based on userId, and ingredientId
	filePath := filepath.Join(fmt.Sprintf("images/%s/ingredients/%s/%s.jpg", userId, ingredientId, ingredientStockId))

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return "", err
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write the byte slice to the file
	if _, err := file.Write(imgBytes); err != nil {
		return "", err
	}

	// Return the file path as the image URL
	return filePath, nil
}
