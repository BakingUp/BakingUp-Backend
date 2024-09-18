package service

import (
	"fmt"
	"sort"

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

func (s *RecipeService) GetRecipeDetail(c *fiber.Ctx, recipeID string) (*domain.RecipeDetail, error) {
	recipe, err := s.recipeRepo.GetRecipeDetail(c, recipeID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, recipe.UserID)
	if err != nil {
		return nil, err
	}

	var recipeIngredients []domain.RecipeIngredient
	for _, recipeIngredientItem := range recipe.RecipeIngredients() {
		firstIngredientURL := recipeIngredientItem.Ingredient().IngredientImages()[0].IngredientURL
		recipeIngredient := &domain.RecipeIngredient{
			IngredientName:     util.GetIngredientName(recipeIngredientItem.Ingredient(), language),
			IngredientURL:      firstIngredientURL,
			IngredientQuantity: util.CombineIngredientQuantity(recipeIngredientItem.RecipeIngredientQuantity, recipeIngredientItem.Ingredient().Unit),
			IngredientPrice:    util.CalculateIngredientPrice(recipeIngredientItem.Ingredient().IngredientDetail()),
		}

		recipeIngredients = append(recipeIngredients, *recipeIngredient)
	}

	totalTimeHours := recipe.TotalTime.Hour()
	totalTimeMinutes := recipe.TotalTime.Minute()

	var recipeURLs []string
	recipeImages := recipe.RecipeImages()

	sort.Slice(recipeImages, func(i, j int) bool {
		return recipeImages[i].ImageIndex < recipeImages[j].ImageIndex
	})
	for _, image := range recipeImages {
		recipeURLs = append(recipeURLs, image.RecipeURL)
	}

	var instructionURLs []string
	instructionImages := recipe.RecipeInstructionImages()
	sort.Slice(instructionImages, func(i, j int) bool {
		return instructionImages[i].InstructionImageIndex < instructionImages[j].InstructionImageIndex
	})

	for _, image := range instructionImages {
		instructionURLs = append(instructionURLs, image.InstructionURL)
	}

	orderAmount := 0
	for _, orderProductItem := range recipe.OrderProducts() {
		if orderProductItem.RecipeID == recipe.RecipeID {
			orderAmount += orderProductItem.ProductQuantity
		}
	}

	recipeDetail := &domain.RecipeDetail{
		Status:            1,
		RecipeName:        util.GetRecipeName(recipe, language),
		RecipeURL:         recipeURLs,
		TotalTime:         fmt.Sprintf("%d hr %d mins", totalTimeHours, totalTimeMinutes),
		Servings:          recipe.Serving,
		Stars:             4,
		NumOfOrder:        orderAmount,
		RecipeIngredients: recipeIngredients,
		InstructionURL:    instructionURLs,
		InstructionSteps:  util.GetInstructionSteps(recipe, language),
	}

	return recipeDetail, nil
}