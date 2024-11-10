package repository

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
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
		db.Recipes.RecipeInstructionImages.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (rr *RecipeRepository) DeleteRecipe(c *fiber.Ctx, recipeID string) error {
	_, err := rr.db.Recipes.FindMany(
		db.Recipes.RecipeID.Equals(recipeID),
	).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) AddRecipe(c *fiber.Ctx, recipe *domain.AddRecipePayload) error {
	_, err := rr.db.Recipes.CreateOne(
		db.Recipes.RecipeID.Set(recipe.RecipeID),
		db.Recipes.User.Link(
			db.Users.UserID.Equals(recipe.UserID),
		),
		db.Recipes.RecipeEngName.Set(recipe.RecipeEngName),
		db.Recipes.RecipeThaiName.Set(recipe.RecipeThaiName),
		db.Recipes.TotalTime.Set(recipe.TotalTime),
		db.Recipes.Serving.Set(recipe.Servings),
		db.Recipes.EngInstruction.Set(recipe.EngInstruction),
		db.Recipes.ThaiInstruction.Set(recipe.ThaiInstruction),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) AddRecipeImage(c *fiber.Ctx, recipeImage *domain.AddRecipeImagePayload) error {
	_, err := rr.db.RecipeImages.CreateOne(
		db.RecipeImages.Recipe.Link(
			db.Recipes.RecipeID.Equals(recipeImage.RecipeID),
		),
		db.RecipeImages.ImageIndex.Set(recipeImage.ImageIndex),
		db.RecipeImages.RecipeURL.Set(recipeImage.ImgUrl),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) AddRecipeInstructionImage(c *fiber.Ctx, recipeInstructionImage *domain.AddRecipeInstructionImagePayload) error {
	_, err := rr.db.RecipeInstructionImages.CreateOne(
		db.RecipeInstructionImages.Recipe.Link(
			db.Recipes.RecipeID.Equals(recipeInstructionImage.RecipeID),
		),
		db.RecipeInstructionImages.InstructionImageIndex.Set(recipeInstructionImage.ImageIndex),
		db.RecipeInstructionImages.InstructionURL.Set(recipeInstructionImage.ImgUrl),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) AddRecipeIngredient(c *fiber.Ctx, recipeIngredient *domain.AddRecipeIngredientPayload) error {
	_, err := rr.db.RecipeIngredients.CreateOne(
		db.RecipeIngredients.Recipe.Link(
			db.Recipes.RecipeID.Equals(recipeIngredient.RecipeID),
		),
		db.RecipeIngredients.Ingredient.Link(
			db.Ingredients.IngredientID.Equals(recipeIngredient.IngredientID),
		),
		db.RecipeIngredients.RecipeIngredientQuantity.Set(recipeIngredient.RecipeIngredientQuantity),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) UpdateHiddenCost(c *fiber.Ctx, hiddenCost *domain.UpdateHiddenCostPayload) error {
	_, err := rr.db.Recipes.FindUnique(
		db.Recipes.RecipeID.Equals(hiddenCost.RecipeID),
	).Update(
		db.Recipes.HiddenCost.Set(hiddenCost.HiddenCost),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) UpdateLaborCost(c *fiber.Ctx, laborCost *domain.UpdateLaborCostPayload) error {
	_, err := rr.db.Recipes.FindUnique(
		db.Recipes.RecipeID.Equals(laborCost.RecipeID),
	).Update(
		db.Recipes.LaborCost.Set(laborCost.LaborCost),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) UpdateProfitMargin(c *fiber.Ctx, profitMargin *domain.UpdateProfitMarginPayload) error {
	_, err := rr.db.Recipes.FindUnique(
		db.Recipes.RecipeID.Equals(profitMargin.RecipeID),
	).Update(
		db.Recipes.ProfitMargin.Set(profitMargin.ProfitMargin),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) EditRecipe(c *fiber.Ctx, recipe *domain.EditRecipePayload) error {
	_, err := rr.db.Recipes.FindUnique(
		db.Recipes.RecipeID.Equals(recipe.RecipeID),
	).Update(
		db.Recipes.RecipeEngName.Set(recipe.RecipeEngName),
		db.Recipes.RecipeThaiName.Set(recipe.RecipeThaiName),
		db.Recipes.TotalTime.Set(recipe.TotalTime),
		db.Recipes.Serving.Set(recipe.Servings),
		db.Recipes.EngInstruction.Set(recipe.EngInstruction),
		db.Recipes.ThaiInstruction.Set(recipe.ThaiInstruction),
	).Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) DeleteRecipeIngredients(c *fiber.Ctx, recipeID string) error {
	_, err := rr.db.RecipeIngredients.FindMany(
		db.RecipeIngredients.RecipeID.Equals(recipeID),
	).Delete().Exec(c.Context())

	if err != nil {
		return err
	}

	return nil
}

func (rr *RecipeRepository) GetEditRecipeDetail(c *fiber.Ctx, recipeID string) (*db.RecipesModel, error) {
	recipe, err := rr.db.Recipes.FindUnique(
		db.Recipes.RecipeID.Equals(recipeID),
	).With(
		db.Recipes.RecipeIngredients.Fetch().With(
			db.RecipeIngredients.Ingredient.Fetch().With(
				db.Ingredients.IngredientImages.Fetch(),
				db.Ingredients.IngredientDetail.Fetch(),
			),
		),
		db.Recipes.RecipeInstructionImages.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (rr *RecipeRepository) GetRecipeStarData(c *fiber.Ctx, userID string) ([]db.OrdersModel, error) {
	orders, err := rr.db.Orders.FindMany(
		db.Orders.UserID.Equals(userID),
	).With(
		db.Orders.OrderProducts.Fetch().With(db.OrderProducts.Recipe.Fetch().With(db.Recipes.Stocks.Fetch().With(db.Stocks.StockDetail.Fetch()), db.Recipes.RecipeImages.Fetch())),
		db.Orders.User.Fetch().With(db.Users.FixCosts.Fetch()),
		db.Orders.CuttingStock.Fetch(),
	).Exec(c.Context())

	if err != nil {
		return nil, err
	}

	return orders, nil
}
