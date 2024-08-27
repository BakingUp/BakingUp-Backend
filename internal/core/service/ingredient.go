package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type IngredientService struct {
	ingredientRepo port.IngredientRepository
}

func NewIngredientService(ingredientRepo port.IngredientRepository) *IngredientService {
	return &IngredientService{
		ingredientRepo: ingredientRepo,
	}
}

func (s *IngredientService) GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*domain.IngredientDetail, error) {
	ingredient, err := s.ingredientRepo.GetIngredientDetail(c, ingredientID)
	
	if err != nil {
		return nil, err
	}

	totalQuantity := 0.0
	var stocks []domain.Stock
	for _, detail := range ingredient.IngredientDetail() {
		totalQuantity += detail.IngredientQuantity
		stocks = append(stocks, domain.Stock{
			StockID:          detail.IngredientStockID,
			StockURL:         detail.IngredientStockURL,
			Price:            util.CombinePrice(detail.Price, ingredient.Unit),
			Quantity:         util.CombineIngredientQuantity(detail.IngredientQuantity, ingredient.Unit),
			ExpirationDate:   detail.ExpirationDate.Format("02/01/2006"),
			ExpirationStatus: util.CalculateExpirationStatus(detail.ExpirationDate),
		})

	}

	var ingredientURLs []string
	for _, image := range ingredient.IngredientImages() {
		ingredientURLs = append(ingredientURLs, image.IngredientURL)
	}

	detail := &domain.IngredientDetail{
		IngredientName:     ingredient.IngredientEngName,
		IngredientQuantity: util.CombineIngredientQuantity(totalQuantity, ingredient.Unit),
		StockAmount:        len(stocks),
		IngredientURL:      ingredientURLs,
		IngredientLessThan: int(ingredient.IngredientLessThan),
		Stocks:             stocks,
	}

	return detail, nil
}
