package util

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
)

func GetRecipeName(recipe *db.RecipesModel, language *db.Language) string {
	if *language == db.LanguageTh {
		return recipe.RecipeThaiName
	}

	return recipe.RecipeEngName
}
