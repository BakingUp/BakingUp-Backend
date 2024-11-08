package port

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type RecipeRepository interface {
	GetAllRecipes(c *fiber.Ctx, userID string) ([]db.RecipesModel, error)
	GetRecipeDetail(c *fiber.Ctx, recipeID string) (*db.RecipesModel, error)
	DeleteRecipe(c *fiber.Ctx, recipeID string) error
	AddRecipe(c *fiber.Ctx, recipe *domain.AddRecipePayload) error
	AddRecipeImage(c *fiber.Ctx, recipeImage *domain.AddRecipeImagePayload) error
	AddRecipeInstructionImage(c *fiber.Ctx, recipeInstructionImage *domain.AddRecipeInstructionImagePayload) error
	AddRecipeIngredient(c *fiber.Ctx, recipeIngredient *domain.AddRecipeIngredientPayload) error
	UpdateHiddenCost(c *fiber.Ctx, hiddenCost *domain.UpdateHiddenCostPayload) error
	UpdateLaborCost(c *fiber.Ctx, laborCost *domain.UpdateLaborCostPayload) error
	UpdateProfitMargin(c *fiber.Ctx, profitMargin *domain.UpdateProfitMarginPayload) error
	EditRecipe(c *fiber.Ctx, recipe *domain.EditRecipePayload) error
	DeleteRecipeIngredients(c *fiber.Ctx, recipeID string) error
	GetEditRecipeDetail(c *fiber.Ctx, recipeID string) (*db.RecipesModel, error)
}

type RecipeService interface {
	GetAllRecipes(c *fiber.Ctx, userID string) (*domain.RecipeList, error)
	GetRecipeDetail(c *fiber.Ctx, recipeID string) (*domain.RecipeDetail, error)
	DeleteRecipe(c *fiber.Ctx, recipeID string) error
	AddRecipe(c *fiber.Ctx, recipe *domain.AddRecipeRequest) error
	UpdateHiddenCost(c *fiber.Ctx, hiddenCost *domain.UpdateHiddenCostRequest) error
	UpdateLaborCost(c *fiber.Ctx, laborCost *domain.UpdateLaborCostRequest) error
	UpdateProfitMargin(c *fiber.Ctx, profitMargin *domain.UpdateProfitMarginRequest) error
	EditRecipe(c *fiber.Ctx, recipe *domain.EditRecipeRequest) error
	GetEditRecipeDetail(c *fiber.Ctx, recipeID string) (*domain.GetEditRecipeDetail, error)
}
