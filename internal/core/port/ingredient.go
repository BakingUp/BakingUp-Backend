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
}

type IngredientService interface {
	GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*domain.IngredientDetail, error)
	GetAllIngredients(c *fiber.Ctx, userID string) (*domain.IngredientList, error)
	GetIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*domain.IngredientStockDetail, error)
}
