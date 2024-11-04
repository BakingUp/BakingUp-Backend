package util

import (
	"encoding/base64"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/BakingUp/BakingUp-Backend/prisma/db"
)

func GetRecipeName(recipe *db.RecipesModel, language *db.Language) string {
	if *language == db.LanguageTh {
		return recipe.RecipeThaiName
	}

	return recipe.RecipeEngName
}

func GetInstructionSteps(recipe *db.RecipesModel, language *db.Language) string {
	if *language == db.LanguageTh {
		return recipe.ThaiInstruction
	}

	return recipe.EngInstruction
}

func CalculateIngredientPrice(value []db.IngredientDetailModel, recipeIngredientQuantity float64) float64 {
	var price float64

	if len(value) == 0 {
		return -1
	}

	for _, detail := range value {
		if time.Now().Before(detail.ExpirationDate) {
			price += detail.Price / detail.IngredientQuantity
		}
	}

	finalPrice := price / float64(len(value))
	return math.Round(finalPrice*recipeIngredientQuantity*100) / 100
}

func FormatTotalTime(totalTimeHours, totalTimeMinutes int) string {
	var result string

	if totalTimeHours > 0 {
		result += fmt.Sprintf("%d %s", totalTimeHours, func() string {
			if totalTimeHours == 1 {
				return "hr"
			}
			return "hrs"
		}())
	}

	if totalTimeMinutes > 0 {
		if result != "" {
			result += " "
		}
		result += fmt.Sprintf("%d %s", totalTimeMinutes, func() string {
			if totalTimeMinutes == 1 {
				return "min"
			}
			return "mins"
		}())
	}

	return result
}

func UploadRecipeImage(userId string, recipeId string, imgBase64 string, index string) (string, error) {
	// Decode the base64 string to a byte slice
	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", err
	}
	// Create the path based on userId, and ingredientId
	filePath := filepath.Join(fmt.Sprintf("images/%s/recipes/%s/recipeImgs/%s.jpg", userId, recipeId, index))

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

func UploadInstructionImage(userId string, recipeId string, imgBase64 string, index string) (string, error) {
	// Decode the base64 string to a byte slice
	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", err
	}
	// Create the path based on userId, and ingredientId
	filePath := filepath.Join(fmt.Sprintf("images/%s/recipes/%s/instructionImgs/%s.jpg", userId, recipeId, index))

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

func ConvertToTime(hours, mins string) (time.Time, error) {
	const layout = "2 Jan 2006 15:04:05"
	baseTime, err := time.Parse(layout, "1 Jan 2000 00:00:00")
	if err != nil {
		return time.Time{}, err
	}

	hoursDuration, err := time.ParseDuration(hours + "h")
	if err != nil {
		return time.Time{}, err
	}

	minsDuration, err := time.ParseDuration(mins + "m")
	if err != nil {
		return time.Time{}, err
	}

	totalTime := baseTime.Add(hoursDuration).Add(minsDuration)
	return totalTime, nil
}
