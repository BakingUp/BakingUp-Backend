package util

import (
	"math"
	"sort"
	"time"

	"github.com/BakingUp/BakingUp-Backend/prisma/db"
)

func GetRecipeName(recipe *db.RecipesModel, language *db.Language) string {
	if *language == db.LanguageTh {
		return recipe.RecipeThaiName
	}

	return recipe.RecipeEngName
}

func GetInstructionSteps(recipe *db.RecipesModel, language *db.Language) []string {
	var urls []string

	if *language == db.LanguageTh {
		steps := recipe.RecipeThaiInstructionSteps()

		sort.Slice(steps, func(i, j int) bool {
			return steps[i].InstructionOrder < steps[j].InstructionOrder
		})

		for _, step := range steps {
			urls = append(urls, step.InstructionStep)
		}
	} else {
		steps := recipe.RecipeEngInstructionSteps()

		sort.Slice(steps, func(i, j int) bool {
			return steps[i].InstructionOrder < steps[j].InstructionOrder
		})

		for _, step := range steps {
			urls = append(urls, step.InstructionStep)
		}
	}

	return urls
}

func CalculateIngredientPrice(value []db.IngredientDetailModel) float64 {
	var price float64
	var totalQuantity float64

	for _, detail := range value {
		if time.Now().After(detail.ExpirationDate) {
			price += detail.IngredientQuantity * detail.Price
			totalQuantity += detail.IngredientQuantity
		}
	}

	if totalQuantity == 0 || price == 0 {
		return -1
	}

	finalPrice := price / totalQuantity
	return math.Round(finalPrice * 100) / 100
}
