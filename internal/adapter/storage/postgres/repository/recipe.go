package repository

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type RecipeRepository struct {
	db *db.PrismaClient
}

func NewRecipeRepository(db *db.PrismaClient) *RecipeRepository {
	return &RecipeRepository{
		db,
	}
}

func (rr *RecipeRepository) GetAllRecipes(c *fiber.Ctx, userID string) ([]db.RecipesModel, error) {
	recipes, err := rr.db.Recipes.FindMany(
		db.Recipes.UserID.Equals(userID),
	).With(
		db.Recipes.OrderProducts.Fetch(),
		db.Recipes.RecipeImages.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (rr *RecipeRepository) GetRecipeDetail(c *fiber.Ctx, recipeID string) (*db.RecipesModel, error) {
	recipe, err := rr.db.Recipes.FindUnique(
		db.Recipes.RecipeID.Equals(recipeID),
	).With(
		db.Recipes.OrderProducts.Fetch(),
		db.Recipes.RecipeImages.Fetch(),
		db.Recipes.RecipeIngredients.Fetch().With(
			db.RecipeIngredients.Ingredient.Fetch().With(
				db.Ingredients.IngredientImages.Fetch(),
				db.Ingredients.IngredientDetail.Fetch(),
			),
		),
		db.Recipes.RecipeEngInstructionSteps.Fetch(),
		db.Recipes.RecipeThaiInstructionSteps.Fetch(),
		db.Recipes.RecipeInstructionImages.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return recipe, nil
}