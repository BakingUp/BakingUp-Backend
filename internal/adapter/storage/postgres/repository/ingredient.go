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

func (ir *IngredientRepository) GetIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*db.IngredientDetailModel, error) {
	ingredient, err := ir.db.IngredientDetail.FindFirst(
		db.IngredientDetail.IngredientStockID.Equals(ingredientStockID),
	).With(
		db.IngredientDetail.Ingredient.Fetch(),
		db.IngredientDetail.IngredientNotes.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (ir *IngredientRepository) DeleteIngredientBatchNote(c *fiber.Ctx, ingredientNoteID string) error {
	_, err := ir.db.IngredientNotes.FindMany(
		db.IngredientNotes.IngredientNoteID.Equals(ingredientNoteID),
	).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}