package repository

import (
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type IngredientRepository struct {
	db *db.PrismaClient
}

func NewIngredientRepository(db *db.PrismaClient) *IngredientRepository {
	return &IngredientRepository{
		db,
	}
}

func (ir *IngredientRepository) GetAllIngredients(c *fiber.Ctx, userID string) ([]db.IngredientsModel, error) {
	ingredients, err := ir.db.Ingredients.FindMany(
		db.Ingredients.UserID.Equals(userID),
	).With(
		db.Ingredients.IngredientDetail.Fetch(),
		db.Ingredients.IngredientImages.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return ingredients, err
}

func (ir *IngredientRepository) GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*db.IngredientsModel, error) {
	ingredient, err := ir.db.Ingredients.FindFirst(
		db.Ingredients.IngredientID.Equals(ingredientID),
	).With(
		db.Ingredients.IngredientDetail.Fetch(),
		db.Ingredients.IngredientImages.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return ingredient, nil
}
