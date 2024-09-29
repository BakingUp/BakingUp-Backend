package service

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type IngredientService struct {
	ingredientRepo port.IngredientRepository
	userService    port.UserService
}

func NewIngredientService(ingredientRepo port.IngredientRepository, userService port.UserService) *IngredientService {
	return &IngredientService{
		ingredientRepo: ingredientRepo,
		userService:    userService,
	}
}

func (s *IngredientService) GetAllIngredients(c *fiber.Ctx, userID string) (*domain.IngredientList, error) {
	ingredients, err := s.ingredientRepo.GetAllIngredients(c, userID)

	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, userID)

	if err != nil {
		return nil, err
	}

	expirationDate, err := s.userService.GetUserExpirationDate(c, userID)

	if err != nil {
		return nil, err
	}

	var ingredientItems []domain.Ingredient
	for _, item := range ingredients {
		var stockAmount int
		var stockQuantity int
		var stockExpirationDate time.Time
		var IImage string
		for _, ingredientDetailItem := range item.IngredientDetail() {
			stockQuantity += int(ingredientDetailItem.IngredientQuantity)
			if stockAmount == 0 {
				stockExpirationDate = ingredientDetailItem.ExpirationDate
			} else if ingredientDetailItem.ExpirationDate.Before(stockExpirationDate) {
				stockExpirationDate = ingredientDetailItem.ExpirationDate
			}
			stockAmount++
		}

		for _, ingredientImageItem := range item.IngredientImages() {
			if ingredientImageItem.IngredientID == item.IngredientID {
				IImage = ingredientImageItem.IngredientURL
				break
			}
		}

		unit := string(item.Unit)
		unit = strings.ToLower(unit)
		ingredientItem := &domain.Ingredient{
			IngredientId:     item.IngredientID,
			IngredientName:   util.GetIngredientName(&item, language),
			Quantity:         fmt.Sprintf("%d %s", stockQuantity, unit),
			Stock:            stockAmount,
			IngredientURL:    IImage,
			ExpirationStatus: util.CalculateExpirationStatus(stockExpirationDate, expirationDate.BlackExpirationDate, expirationDate.RedExpirationDate, expirationDate.YellowExpirationDate),
		}

		ingredientItems = append(ingredientItems, *ingredientItem)
	}

	ingredientList := &domain.IngredientList{
		Ingredients: ingredientItems,
	}

	return ingredientList, nil

}

func (s *IngredientService) GetIngredientDetail(c *fiber.Ctx, ingredientID string) (*domain.IngredientDetail, error) {
	ingredient, err := s.ingredientRepo.GetIngredientDetail(c, ingredientID)

	if err != nil {
		return nil, err
	}

	language, err := s.userService.GetUserLanguage(c, ingredient.UserID)
	if err != nil {
		return nil, err
	}

	expirationDate, err := s.userService.GetUserExpirationDate(c, ingredient.UserID)

	if err != nil {
		return nil, err
	}

	totalQuantity := 0.0
	var stocks []domain.Stock
	for _, detail := range ingredient.IngredientDetail() {
		ingredientStockURL, _ := detail.IngredientStockURL()
		totalQuantity += detail.IngredientQuantity
		stocks = append(stocks, domain.Stock{
			StockID:          detail.IngredientStockID,
			StockURL:         ingredientStockURL,
			Price:            util.CombinePrice(detail.Price, ingredient.Unit),
			Quantity:         util.CombineIngredientQuantity(detail.IngredientQuantity, ingredient.Unit),
			ExpirationDate:   detail.ExpirationDate.Format("02/01/2006"),
			ExpirationStatus: util.CalculateExpirationStatus(detail.ExpirationDate, expirationDate.BlackExpirationDate, expirationDate.RedExpirationDate, expirationDate.YellowExpirationDate),
		})

	}

	var ingredientURLs []string
	ingredientImages := ingredient.IngredientImages()

	sort.Slice(ingredientImages, func(i, j int) bool {
		return ingredientImages[i].IngredientImageIndex < ingredientImages[j].IngredientImageIndex
	})
	for _, image := range ingredient.IngredientImages() {
		ingredientURLs = append(ingredientURLs, image.IngredientURL)
	}

	detail := &domain.IngredientDetail{
		IngredientName:     util.GetIngredientName(ingredient, language),
		IngredientQuantity: util.CombineIngredientQuantity(totalQuantity, ingredient.Unit),
		StockAmount:        len(stocks),
		IngredientURL:      ingredientURLs,
		IngredientLessThan: int(ingredient.IngredientLessThan),
		Stocks:             stocks,
	}

	return detail, nil
}

func (s *IngredientService) GetIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*domain.IngredientStockDetail, error) {
	ingredient, err := s.ingredientRepo.GetIngredientStockDetail(c, ingredientStockID)

	if err != nil {
		return nil, err
	}

	var notes []domain.IngredientNote
	ingredientNotes := ingredient.IngredientNotes()

	sort.SliceStable(ingredientNotes, func(i, j int) bool {
		return ingredientNotes[i].NoteCreatedAt.After(ingredientNotes[j].NoteCreatedAt)
	})

	for _, note := range ingredientNotes {
		notes = append(notes, domain.IngredientNote{
			IngredientNoteID: note.IngredientNoteID,
			IngredientNote:   note.IngredientNote,
			NoteCreatedAt:    note.NoteCreatedAt.Format("02/01/2006"),
		})
	}
	ingredientStockURL, _ := ingredient.IngredientStockURL()
	stockDetail := &domain.IngredientStockDetail{
		IngredientEngName:  ingredient.Ingredient().IngredientEngName,
		IngredientThaiName: ingredient.Ingredient().IngredientThaiName,
		IngredientQuantity: util.CombineIngredientQuantity(ingredient.IngredientQuantity, ingredient.Ingredient().Unit),
		IngredientPrice:    strconv.FormatFloat(ingredient.Price, 'f', -1, 64),
		IngredientBrand:    ingredient.IngredientBrand,
		IngredientSupplier: ingredient.IngredientSupplier,
		IngredientStockURL: ingredientStockURL,
		DayBeforeExpire:    ingredient.ExpirationDate.Format("02/01/2006"),
		Notes:              notes,
	}

	return stockDetail, nil
}

func (s *IngredientService) DeleteIngredientBatchNote(c *fiber.Ctx, ingredientNoteID string) error {
	err := s.ingredientRepo.DeleteIngredientBatchNote(c, ingredientNoteID)
	if err != nil {
		return err
	}

	return nil
}

func (s *IngredientService) DeleteIngredient(c *fiber.Ctx, ingredientID string) error {
	err := s.ingredientRepo.DeleteIngredient(c, ingredientID)
	if err != nil {
		return err
	}

	return nil
}
