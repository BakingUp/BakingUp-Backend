package repository

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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

func (ir *IngredientRepository) GetAddEditIngredientStockDetail(c *fiber.Ctx, ingredientID string) (*db.IngredientsModel, error) {
	ingredient, err := ir.db.Ingredients.FindFirst(
		db.Ingredients.IngredientID.Equals(ingredientID),
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

func (ir *IngredientRepository) DeleteIngredient(c *fiber.Ctx, ingredientID string) error {
	_, err := ir.db.Ingredients.FindMany(
		db.Ingredients.IngredientID.Equals(ingredientID),
	).Delete().Exec(c.Context())
	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) DeleteIngredientStock(c *fiber.Ctx, ingredientStockID string) error {
	_, err := ir.db.IngredientDetail.FindMany(
		db.IngredientDetail.IngredientStockID.Equals(ingredientStockID),
	).Delete().Exec(c.Context())
	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) AddIngredient(c *fiber.Ctx, ingredient *domain.AddIngredientPayload) error {
	_, err := ir.db.Ingredients.CreateOne(
		db.Ingredients.IngredientID.Set(ingredient.IngredientID),
		db.Ingredients.User.Link(
			db.Users.UserID.Equals(ingredient.UserID),
		),
		db.Ingredients.IngredientEngName.Set(ingredient.IngredientEngName),
		db.Ingredients.IngredientThaiName.Set(ingredient.IngredientThaiName),
		db.Ingredients.Unit.Set(db.Unit(ingredient.Unit)),
		db.Ingredients.IngredientLessThan.Set(float64(ingredient.StockLessThan)),
		db.Ingredients.DayBeforeExpire.Set(ingredient.DayBeforeExpire),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) AddIngredientImage(c *fiber.Ctx, ingredientImage *domain.AddIngredientImagePayload) error {
	_, err := ir.db.IngredientImages.CreateOne(
		db.IngredientImages.IngredientImageIndex.Set(ingredientImage.ImageIndex),
		db.IngredientImages.Ingredient.Link(
			db.Ingredients.IngredientID.Equals(ingredientImage.IngredientID),
		),
		db.IngredientImages.IngredientURL.Set(ingredientImage.ImgUrl),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) AddIngredientStock(c *fiber.Ctx, ingredientStock *domain.AddIngredientStockPayload) error {
	_, err := ir.db.IngredientDetail.CreateOne(
		db.IngredientDetail.IngredientStockID.Set(ingredientStock.IngredientStockID),
		db.IngredientDetail.Ingredient.Link(
			db.Ingredients.IngredientID.Equals(ingredientStock.IngredientID),
		),
		db.IngredientDetail.Price.Set(ingredientStock.Price),
		db.IngredientDetail.IngredientQuantity.Set(ingredientStock.IngredientQuantity),
		db.IngredientDetail.ExpirationDate.Set(ingredientStock.ExpirationDate),
		db.IngredientDetail.IngredientSupplier.Set(ingredientStock.IngredientSupplier),
		db.IngredientDetail.IngredientBrand.Set(ingredientStock.IngredientBrand),
		db.IngredientDetail.IngredientStockURL.Set(ingredientStock.IngredientStockURL),
	).Exec(c.Context())
	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) AddIngredientNote(c *fiber.Ctx, ingredientNote *domain.AddIngredientNotePayload) error {
	_, err := ir.db.IngredientNotes.CreateOne(
		db.IngredientNotes.IngredientNoteID.Set(ingredientNote.IngredientNoteID),
		db.IngredientNotes.IngredientDetail.Link(
			db.IngredientDetail.IngredientStockID.Equals(ingredientNote.IngredientStockID),
		),
		db.IngredientNotes.NoteCreatedAt.Set(ingredientNote.NoteCreatedAt),
		db.IngredientNotes.IngredientNote.Set(ingredientNote.Note),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) GetUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientID string) (float64, error) {
	ingredient, err := ir.db.Ingredients.FindFirst(
		db.Ingredients.IngredientID.Equals(ingredientID),
	).Exec(c.Context())

	if err != nil {
		return 0, err
	}

	return ingredient.IngredientLessThan, nil
}

func (ir *IngredientRepository) DeleteUnexpiredIngredient(c *fiber.Ctx, ingredientStockID string) error {
	_, err := ir.db.IngredientDetail.FindMany(
		db.IngredientDetail.IngredientStockID.Equals(ingredientStockID),
	).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) UpdateUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientStockID string, quantity float64) error {
	_, err := ir.db.IngredientDetail.FindUnique(
		db.IngredientDetail.IngredientStockID.Equals(ingredientStockID),
	).Update(
		db.IngredientDetail.IngredientQuantity.Set(quantity),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) EditIngredient(c *fiber.Ctx, ingredient *domain.EditIngredientPayload) error {
	_, err := ir.db.Ingredients.FindUnique(
		db.Ingredients.IngredientID.Equals(ingredient.IngredientID),
	).Update(
		db.Ingredients.IngredientEngName.Set(ingredient.IngredientEngName),
		db.Ingredients.IngredientThaiName.Set(ingredient.IngredientThaiName),
		db.Ingredients.IngredientLessThan.Set(float64(ingredient.StockLessThan)),
		db.Ingredients.DayBeforeExpire.Set(ingredient.DayBeforeExpire),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) GetAddEditIngredientDetail(c *fiber.Ctx, ingredientID string) (*db.IngredientsModel, error) {
	ingredient, err := ir.db.Ingredients.FindFirst(
		db.Ingredients.IngredientID.Equals(ingredientID),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (ir *IngredientRepository) EditIngredientStock(c *fiber.Ctx, ingredientStock *domain.EditIngredientStockPayload) error {
	_, err := ir.db.IngredientDetail.FindUnique(
		db.IngredientDetail.IngredientStockID.Equals(ingredientStock.IngredientStockID),
	).Update(
		db.IngredientDetail.Price.Set(ingredientStock.Price),
		db.IngredientDetail.IngredientQuantity.Set(ingredientStock.Quantity),
		db.IngredientDetail.ExpirationDate.Set(ingredientStock.ExpirationDate),
		db.IngredientDetail.IngredientSupplier.Set(ingredientStock.Supplier),
		db.IngredientDetail.IngredientBrand.Set(ingredientStock.Brand),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (ir *IngredientRepository) GetEditIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*db.IngredientDetailModel, error) {
	ingredient, err := ir.db.IngredientDetail.FindFirst(
		db.IngredientDetail.IngredientStockID.Equals(ingredientStockID),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return ingredient, nil
}