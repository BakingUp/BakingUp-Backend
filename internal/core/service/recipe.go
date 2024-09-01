package service

import (
	"fmt"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type RecipeService struct {
	recipeRepo  port.RecipeRepository
	userService port.UserService
}

func NewRecipeService(recipeRepo port.RecipeRepository, userService port.UserService) *RecipeService {
	return &RecipeService{
		recipeRepo:  recipeRepo,
		userService: userService,
	}
}

func (s *RecipeService) GetAllRecipes(c *fiber.Ctx, userID string) (*domain.RecipeList, error) {
	recipes, err := s.recipeRepo.GetAllRecipes(c, userID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}
	var recipeItems []domain.Recipe
	for _, item := range recipes {
		var orderAmount int
		var recipeImg string

		for _, orderProductItem := range item.OrderProducts() {
			if orderProductItem.RecipeID == item.RecipeID {
				orderAmount += orderProductItem.ProductQuantity
			}
		}

		for _, recipeImageItem := range item.RecipeImages() {
			if recipeImageItem.RecipeID == item.RecipeID {
				recipeImg = recipeImageItem.RecipeURL
			}
		}
		totalTimeHours := item.TotalTime.Hour()
		totalTimeMinutes := item.TotalTime.Minute()

		recipeItem := &domain.Recipe{
			RecipeName: util.GetRecipeName(&item, language),
			RecipeImg:  recipeImg,
			TotalTime:  fmt.Sprintf("%d hr %d mins", totalTimeHours, totalTimeMinutes),
			Servings:   item.Serving,
			Stars:      4,
			NumOfOrder: orderAmount,
		}

		recipeItems = append(recipeItems, *recipeItem)
	}

	recipeList := &domain.RecipeList{
		Recipes: recipeItems,
	}

	return recipeList, nil
}
