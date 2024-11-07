package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type IngredientRepository interface {
	GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*db.IngredientsModel, error)
	GetAllIngredients(c *fiber.Ctx, userID string) ([]db.IngredientsModel, error)
	GetIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*db.IngredientDetailModel, error)
	GetAddEditIngredientStockDetail(c *fiber.Ctx, ingredientID string) (*db.IngredientsModel, error)
	DeleteIngredientBatchNote(c *fiber.Ctx, ingredientNoteID string) error
	DeleteIngredient(c *fiber.Ctx, ingredientID string) error
	DeleteIngredientStock(c *fiber.Ctx, ingredientStockID string) error
	AddIngredient(c *fiber.Ctx, ingredients *domain.AddIngredientPayload) error
	AddIngredientImage(c *fiber.Ctx, ingredientImage *domain.AddIngredientImagePayload) error
	AddIngredientStock(c *fiber.Ctx, ingredientStock *domain.AddIngredientStockPayload) error
	AddIngredientNote(c *fiber.Ctx, ingredientNote *domain.AddIngredientNotePayload) error
	GetUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientID string) (float64, error)
}

type IngredientService interface {
	GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*domain.IngredientDetail, error)
	GetAllIngredients(c *fiber.Ctx, userID string) (*domain.IngredientList, error)
	GetIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*domain.IngredientStockDetail, error)
	GetAddEditIngredientStockDetail(c *fiber.Ctx, ingredientID string) (*domain.AddEditIngredientStockDetail, error)
	DeleteIngredientBatchNote(c *fiber.Ctx, ingredientNoteID string) error
	DeleteIngredient(c *fiber.Ctx, ingredientID string) error
	DeleteIngredientStock(c *fiber.Ctx, ingredientStockID string) error
	AddIngredient(c *fiber.Ctx, ingredients *domain.AddIngredientRequest) error
	AddIngredientStock(c *fiber.Ctx, ingredientStock *domain.AddIngredientStockRequest) error
	GetUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientID string) (float64, error)
}
