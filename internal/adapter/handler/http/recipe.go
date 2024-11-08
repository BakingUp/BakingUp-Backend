package http

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type RecipeHandler struct {
	svc port.RecipeService
}

func NewRecipeHandler(svc port.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		svc: svc,
	}
}

// GetAllRecipes godoc
// @Summary      Get all recipes
// @Description  Get all recipes by user ID
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true  "User ID"
// @Success      200  {object}  domain.RecipeList  "Success"
// @Failure      400  {object}  response     "Cannot get all recipes"
// @Router       /recipe/getAllRecipes [get]
func (rh *RecipeHandler) GetAllRecipes(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	recipes, err := rh.svc.GetAllRecipes(c, userID)
	if err != nil {
		handleError(c, 400, "Cannot get all recipes", err.Error())
		return nil
	}

	handleSuccess(c, recipes)
	return nil
}

// GetRecipeDetail godoc
// @Summary      Get recipe details
// @Description  Get recipe details by recipe ID
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        recipe_id  query  string  true  "Recipe ID"
// @Success      200  {object}  domain.RecipeDetail  "Success"
// @Failure      400  {object}  response     "Cannot get recipe detail"
// @Router       /recipe/getRecipeDetail [get]
func (rh *RecipeHandler) GetRecipeDetail(c *fiber.Ctx) error {
	recipeID := c.Query("recipe_id")

	recipe, err := rh.svc.GetRecipeDetail(c, recipeID)
	if err != nil {
		handleError(c, 400, "Cannot get recipe detail", err.Error())
		return nil
	}

	handleSuccess(c, recipe)
	return nil
}

// DeleteRecipe godoc
// @Summary      Delete a recipe
// @Description  Delete a recipe by using recipe id
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        recipe_id  query  string  true  "Recipe ID"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot delete a recipe"
// @Router       /recipe/deleteRecipe [delete]
func (rh *RecipeHandler) DeleteRecipe(c *fiber.Ctx) error {
	recipeID := c.Query("recipe_id")

	err := rh.svc.DeleteRecipe(c, recipeID)
	if err != nil {
		handleError(c, 400, "Cannot delete a recipe", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}

// AddRecipe godoc
// @Summary      Add a recipe
// @Description  Add a recipe
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        AddRecipeRequest  body  domain.AddRecipeRequest  true  "Recipe Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot add a recipe"
// @Router       /recipe/addRecipe [post]
func (rh *RecipeHandler) AddRecipe(c *fiber.Ctx) error {
	var recipe *domain.AddRecipeRequest
	if err := c.BodyParser(&recipe); err != nil {
		handleError(c, 400, "Cannot parse request body", err.Error())
		return nil
	}

	err := rh.svc.AddRecipe(c, recipe)
	if err != nil {
		handleError(c, 400, "Cannot add a recipe", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}

// UpdateHiddenCost godoc
// @Summary      Update hidden cost
// @Description  Update hidden cost
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        UpdateHiddenCostRequest  body  domain.UpdateHiddenCostRequest  true  "Hidden Cost Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot update hidden cost"
// @Router       /recipe/updateHiddenCost [put]
func (rh *RecipeHandler) UpdateHiddenCost(c *fiber.Ctx) error {
	var request *domain.UpdateHiddenCostRequest
	if err := c.BodyParser(&request); err != nil {
		handleError(c, 400, "Cannot parse request body", err.Error())
		return nil
	}

	err := rh.svc.UpdateHiddenCost(c, request)
	if err != nil {
		handleError(c, 400, "Cannot update hidden cost", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}

// UpdateLaborCost godoc
// @Summary      Update labor cost
// @Description  Update labor cost
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        UpdateLaborCostRequest  body  domain.UpdateLaborCostRequest  true  "Labor Cost Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot update labor cost"
// @Router       /recipe/updateLaborCost [put]
func (rh *RecipeHandler) UpdateLaborCost(c *fiber.Ctx) error {
	var request *domain.UpdateLaborCostRequest
	if err := c.BodyParser(&request); err != nil {
		handleError(c, 400, "Cannot parse request body", err.Error())
		return nil
	}

	err := rh.svc.UpdateLaborCost(c, request)
	if err != nil {
		handleError(c, 400, "Cannot update labor cost", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}

// UpdateProfitMargin godoc
// @Summary      Update profit margin
// @Description  Update profit margin
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        UpdateProfitMarginRequest  body  domain.UpdateProfitMarginRequest  true  "Profit Margin Request"
// @Success      200  {object}  response  "Success"
// @Failure      400  {object}  response  "Cannot update profit margin"
// @Router       /recipe/updateProfitMargin [put]
func (rh *RecipeHandler) UpdateProfitMargin(c *fiber.Ctx) error {
	var request *domain.UpdateProfitMarginRequest
	if err := c.BodyParser(&request); err != nil {
		handleError(c, 400, "Cannot parse request body", err.Error())
		return nil
	}

	err := rh.svc.UpdateProfitMargin(c, request)
	if err != nil {
		handleError(c, 400, "Cannot update profit margin", err.Error())
		return nil
	}

	handleSuccess(c, nil)
	return nil
}
