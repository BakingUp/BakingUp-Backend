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
	DeleteUnexpiredIngredient(c *fiber.Ctx, ingredientStockID string) error
	UpdateUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientStockID string, quantity float64) error
	EditIngredient(c *fiber.Ctx, ingredient *domain.EditIngredientPayload) error
	GetAddEditIngredientDetail(c *fiber.Ctx, ingredientID string) (*db.IngredientsModel, error)
	EditIngredientStock(c *fiber.Ctx, ingredientStock *domain.EditIngredientStockPayload) error
	GetEditIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*db.IngredientDetailModel, error)
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
	UpdateUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientID string, quantity float64) error
	EditIngredient(c *fiber.Ctx, ingredient *domain.EditIngredientRequest) error
	GetAddEditIngredientDetail(c *fiber.Ctx, ingredientID string) (*domain.GetAddEditIngredientDetail, error)
	EditIngredientStock(c *fiber.Ctx, ingredientStock *domain.EditIngredientStockRequest) error
	GetEditIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*domain.GetEditIngredientStockDetail, error)
	BeforeExpiredIngredientNotifiation() error
}
