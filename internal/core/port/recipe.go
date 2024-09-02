package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type RecipeRepository interface {
	GetAllRecipes(c *fiber.Ctx, userID string) ([]db.RecipesModel, error)
}

type RecipeService interface {
	GetAllRecipes(c *fiber.Ctx, userID string) (*domain.RecipeList, error)
}
