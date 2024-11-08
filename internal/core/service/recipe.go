package service

import (
	"sort"
	"strconv"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
			RecipeID:   item.RecipeID,
			RecipeName: util.GetRecipeName(&item, language),
			RecipeImg:  recipeImg,
			TotalTime:  util.FormatTotalTime(totalTimeHours, totalTimeMinutes),
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
		images := recipeIngredientItem.Ingredient().IngredientImages()
		firstIngredientURL := ""
		if len(images) != 0 {
			firstIngredientURL = images[0].IngredientURL
		}
		recipeIngredient := &domain.RecipeIngredient{
			IngredientName:     util.GetIngredientName(recipeIngredientItem.Ingredient(), language),
			IngredientURL:      firstIngredientURL,
			IngredientQuantity: util.CombineIngredientQuantity(recipeIngredientItem.RecipeIngredientQuantity, recipeIngredientItem.Ingredient().Unit),
			IngredientPrice:    util.CalculateIngredientPrice(recipeIngredientItem.Ingredient().IngredientDetail(), recipeIngredientItem.RecipeIngredientQuantity),
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
		TotalTime:         util.FormatTotalTime(totalTimeHours, totalTimeMinutes),
		Servings:          recipe.Serving,
		Stars:             4,
		NumOfOrder:        orderAmount,
		RecipeIngredients: recipeIngredients,
		InstructionURL:    instructionURLs,
		InstructionSteps:  util.GetInstructionSteps(recipe, language),
		HiddenCost:        recipe.HiddenCost,
		LaborCost:         recipe.LaborCost,
		ProfitMargin:      recipe.ProfitMargin,
	}

	return recipeDetail, nil
}

func (s *RecipeService) DeleteRecipe(c *fiber.Ctx, recipeID string) error {
	err := s.recipeRepo.DeleteRecipe(c, recipeID)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecipeService) AddRecipe(c *fiber.Ctx, payload *domain.AddRecipeRequest) error {
	recipeID := uuid.NewString()
	servings, _ := strconv.Atoi(payload.Servings)
	totalTime, _ := util.ConvertToTime(payload.TotalHours, payload.TotalMins)
	addRecipePayload := &domain.AddRecipePayload{
		UserID:          payload.UserID,
		RecipeID:        recipeID,
		RecipeEngName:   payload.RecipeEngName,
		RecipeThaiName:  payload.RecipeThaiName,
		TotalTime:       totalTime,
		Servings:        servings,
		EngInstruction:  payload.EngInstruction,
		ThaiInstruction: payload.ThaiInstruction,
	}

	err := s.recipeRepo.AddRecipe(c, addRecipePayload)

	if err != nil {
		return err
	}

	for _, ingredient := range payload.Ingredients {
		quantity, _ := strconv.ParseFloat(ingredient.IngredientQuantity, 64)
		addRecipeIngredientPayload := &domain.AddRecipeIngredientPayload{
			RecipeID:                 recipeID,
			IngredientID:             ingredient.IngredientID,
			RecipeIngredientQuantity: quantity,
		}

		err = s.recipeRepo.AddRecipeIngredient(c, addRecipeIngredientPayload)
		if err != nil {
			return err
		}
	}

	imgIndex := 1

	for _, img := range payload.RecipeImg {
		imgUrl, err := util.UploadRecipeImage(payload.UserID, recipeID, img, strconv.Itoa(imgIndex))
		if err != nil {
			return err
		}

		addRecipeImagePayload := &domain.AddRecipeImagePayload{
			RecipeID:   recipeID,
			ImgUrl:     imgUrl,
			ImageIndex: imgIndex,
		}

		err = s.recipeRepo.AddRecipeImage(c, addRecipeImagePayload)
		if err != nil {
			return err
		}

		imgIndex++
	}

	for _, img := range payload.InstructionImg {
		imgUrl, err := util.UploadInstructionImage(payload.UserID, recipeID, img, strconv.Itoa(imgIndex))
		if err != nil {
			return err
		}

		addRecipeInstructionImagePayload := &domain.AddRecipeInstructionImagePayload{
			RecipeID:   recipeID,
			ImgUrl:     imgUrl,
			ImageIndex: imgIndex,
		}

		err = s.recipeRepo.AddRecipeInstructionImage(c, addRecipeInstructionImagePayload)
		if err != nil {
			return err
		}

		imgIndex++
	}

	return nil
}

func (s *RecipeService) UpdateHiddenCost(c *fiber.Ctx, request *domain.UpdateHiddenCostRequest) error {
	hiddenCost, _ := strconv.ParseFloat(request.HiddenCost, 64)

	payload := &domain.UpdateHiddenCostPayload{
		RecipeID:   request.RecipeID,
		HiddenCost: hiddenCost,
	}

	err := s.recipeRepo.UpdateHiddenCost(c, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecipeService) UpdateLaborCost(c *fiber.Ctx, request *domain.UpdateLaborCostRequest) error {
	laborCost, _ := strconv.ParseFloat(request.LaborCost, 64)

	payload := &domain.UpdateLaborCostPayload{
		RecipeID:  request.RecipeID,
		LaborCost: laborCost,
	}

	err := s.recipeRepo.UpdateLaborCost(c, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecipeService) UpdateProfitMargin(c *fiber.Ctx, request *domain.UpdateProfitMarginRequest) error {
	profitMargin, _ := strconv.ParseFloat(request.ProfitMargin, 64)

	payload := &domain.UpdateProfitMarginPayload{
		RecipeID:     request.RecipeID,
		ProfitMargin: profitMargin,
	}

	err := s.recipeRepo.UpdateProfitMargin(c, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecipeService) EditRecipe(c *fiber.Ctx, request *domain.EditRecipeRequest) error {
	totalTime, _ := util.ConvertToTime(request.TotalHours, request.TotalMins)
	servings, _ := strconv.Atoi(request.Servings)

	payload := &domain.EditRecipePayload{
		RecipeID:        request.RecipeID,
		RecipeEngName:   request.RecipeEngName,
		RecipeThaiName:  request.RecipeThaiName,
		TotalTime:       totalTime,
		Servings:        servings,
		EngInstruction:  request.EngInstruction,
		ThaiInstruction: request.ThaiInstruction,
	}

	err := s.recipeRepo.EditRecipe(c, payload)
	if err != nil {
		return err
	}

	err = s.recipeRepo.DeleteRecipeIngredients(c, request.RecipeID)
	if err != nil {
		return err
	}

	for _, ingredient := range request.Ingredients {
		quantity, _ := strconv.ParseFloat(ingredient.IngredientQuantity, 64)
		addRecipeIngredientPayload := &domain.AddRecipeIngredientPayload{
			RecipeID:                 request.RecipeID,
			IngredientID:             ingredient.IngredientID,
			RecipeIngredientQuantity: quantity,
		}

		err = s.recipeRepo.AddRecipeIngredient(c, addRecipeIngredientPayload)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RecipeService) GetEditRecipeDetail(c *fiber.Ctx, recipeID string) (*domain.GetEditRecipeDetail, error) {
	recipe, err := s.recipeRepo.GetEditRecipeDetail(c, recipeID)
	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, recipe.UserID)

	if err != nil {
		return nil, err
	}

	var recipeIngredients []domain.GetEditRecipeIngredientDetail
	for _, recipeIngredientItem := range recipe.RecipeIngredients() {
		images := recipeIngredientItem.Ingredient().IngredientImages()
		firstIngredientURL := ""
		if len(images) != 0 {
			firstIngredientURL = images[0].IngredientURL
		}
		recipeIngredient := &domain.GetEditRecipeIngredientDetail{
			IngredientID:       recipeIngredientItem.IngredientID,
			IngredientName:     util.GetIngredientName(recipeIngredientItem.Ingredient(), language),
			IngredientURL:      firstIngredientURL,
			IngredientQuantity: util.CombineIngredientQuantity(recipeIngredientItem.RecipeIngredientQuantity, recipeIngredientItem.Ingredient().Unit),
		}

		recipeIngredients = append(recipeIngredients, *recipeIngredient)
	}

	totalTimeHours := recipe.TotalTime.Hour()
	totalTimeMinutes := recipe.TotalTime.Minute()

	recipeDetail := &domain.GetEditRecipeDetail{
		RecipeEngName:   recipe.RecipeEngName,
		RecipeThaiName:  recipe.RecipeThaiName,
		TotalHours:      strconv.Itoa(totalTimeHours),
		TotalMins:       strconv.Itoa(totalTimeMinutes),
		Servings:        strconv.Itoa(recipe.Serving),
		EngInstruction:  recipe.EngInstruction,
		ThaiInstruction: recipe.ThaiInstruction,
		Ingredients:     recipeIngredients,
	}

	return recipeDetail, nil
}
