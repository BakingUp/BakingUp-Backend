package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type IngredientRepository interface {
	GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*db.IngredientsModel, error)
}

type IngredientService interface {
	GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*domain.IngredientDetail, error)
}