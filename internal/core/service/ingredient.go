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
	"github.com/google/uuid"
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

	if ingredientList.Ingredients == nil {
		ingredientList.Ingredients = []domain.Ingredient{}
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
	var stockDetails []domain.Stock
	for _, detail := range ingredient.IngredientDetail() {
		ingredientStockURL, _ := detail.IngredientStockURL()
		totalQuantity += detail.IngredientQuantity
		stockDetails = append(stockDetails, domain.Stock{
			StockID:          detail.IngredientStockID,
			StockURL:         ingredientStockURL,
			Price:            util.CombinePrice(detail.Price, ingredient.Unit, detail.IngredientQuantity),
			Quantity:         util.CombineIngredientQuantity(detail.IngredientQuantity, ingredient.Unit),
			ExpirationDate:   detail.ExpirationDate.Format("02/01/2006"),
			ExpirationStatus: util.CalculateExpirationStatus(detail.ExpirationDate, expirationDate.BlackExpirationDate, expirationDate.RedExpirationDate, expirationDate.YellowExpirationDate),
		})
	}

	sort.Slice(stockDetails, func(i, j int) bool {
		dateI, _ := time.Parse("02/01/2006", stockDetails[i].ExpirationDate)
		dateJ, _ := time.Parse("02/01/2006", stockDetails[j].ExpirationDate)
		return dateI.After(dateJ)
	})

	stocks = append(stocks, stockDetails...)

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

func (s *IngredientService) GetAddEditIngredientStockDetail(c *fiber.Ctx, ingredientID string) (*domain.AddEditIngredientStockDetail, error) {
	ingredient, err := s.ingredientRepo.GetAddEditIngredientStockDetail(c, ingredientID)

	if err != nil {
		return nil, err
	}

	detail := &domain.AddEditIngredientStockDetail{
		IngredientEngName:  ingredient.IngredientEngName,
		IngredientThaiName: ingredient.IngredientThaiName,
		Unit:               string(ingredient.Unit),
	}

	return detail, nil
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

func (s *IngredientService) DeleteIngredientStock(c *fiber.Ctx, ingredientStockID string) error {
	err := s.ingredientRepo.DeleteIngredientStock(c, ingredientStockID)
	if err != nil {
		return err
	}

	return nil
}

func (s *IngredientService) AddIngredient(c *fiber.Ctx, ingredients *domain.AddIngredientRequest) error {
	userID := ingredients.UserID
	ingredientID := uuid.NewString()
	dayBeforeExpire := util.ExpirationDate(ingredients.DayBeforeExpire)
	stockLessThan, _ := strconv.Atoi(ingredients.StockLessThan)

	addIngredientPayload := &domain.AddIngredientPayload{
		UserID:             userID,
		IngredientID:       ingredientID,
		IngredientEngName:  ingredients.IngredientEngName,
		IngredientThaiName: ingredients.IngredientThaiName,
		StockLessThan:      stockLessThan,
		Unit:               ingredients.Unit,
		DayBeforeExpire:    dayBeforeExpire,
	}
	err := s.ingredientRepo.AddIngredient(c, addIngredientPayload)
	imgIndex := 1
	for _, img := range ingredients.Img {
		imgUrl, err := util.UploadIngredientImage(userID, ingredientID, img, strconv.Itoa(imgIndex))
		if err != nil {
			return err
		}
		addIngredientImagePayload := &domain.AddIngredientImagePayload{
			IngredientID: ingredientID,
			ImgUrl:       imgUrl,
			ImageIndex:   strconv.Itoa(imgIndex),
		}
		err = s.ingredientRepo.AddIngredientImage(c, addIngredientImagePayload)
		if err != nil {
			return err
		}
		imgIndex++
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *IngredientService) AddIngredientStock(c *fiber.Ctx, ingredientStock *domain.AddIngredientStockRequest) error {
	ingredientID := ingredientStock.IngredientID
	ingredientStockID := uuid.NewString()
	ingredientNoteID := uuid.NewString()
	price, _ := strconv.ParseFloat(ingredientStock.Price, 64)
	quantity, _ := strconv.ParseFloat(ingredientStock.Quantity, 64)
	expirationDate, _ := time.Parse("02/01/2006", ingredientStock.ExpirationDate)
	noteCreatedAt := time.Now()
	userID := ingredientStock.UserID
	ingredientStockImg := ingredientStock.Img

	commonPayload := domain.AddIngredientStockPayload{
		IngredientStockID:  ingredientStockID,
		IngredientID:       ingredientID,
		IngredientQuantity: quantity,
		Price:              price,
		ExpirationDate:     expirationDate,
		IngredientSupplier: ingredientStock.Supplier,
		IngredientBrand:    ingredientStock.IngredientBrand,
		Note:               ingredientStock.Note,
	}

	if ingredientStockImg != "" {
		ingredientStockURL, err := util.UploadIngredientStockImage(userID, ingredientID, ingredientStockID, ingredientStockImg)
		if err != nil {
			return err
		}
		commonPayload.IngredientStockURL = ingredientStockURL
	}

	ingredientStockPayload := &commonPayload

	err := s.ingredientRepo.AddIngredientStock(c, ingredientStockPayload)
	if err != nil {
		return err
	}

	if ingredientStock.Note != "" {
		ingredientNote := &domain.AddIngredientNotePayload{
			IngredientNoteID:  ingredientNoteID,
			IngredientStockID: ingredientStockID,
			Note:              ingredientStock.Note,
			NoteCreatedAt:     noteCreatedAt,
		}

		err = s.ingredientRepo.AddIngredientNote(c, ingredientNote)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *IngredientService) GetUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientID string) (float64, error) {
	ingredient, err := s.ingredientRepo.GetIngredientDetail(c, ingredientID)
	if err != nil {
		return 0, err
	}

	var totalQuantity float64
	for _, detail := range ingredient.IngredientDetail() {
		if detail.ExpirationDate.After(time.Now()) {
			totalQuantity += detail.IngredientQuantity
		}
	}

	return totalQuantity, nil
}

func (s *IngredientService) UpdateUnexpiredIngredientQuantity(c *fiber.Ctx, ingredientID string, quantity float64) error {
	ingredient, err := s.ingredientRepo.GetIngredientDetail(c, ingredientID)
	if err != nil {
		return err
	}

	// Sort ingredient details by expiration date in chronological order
	sort.Slice(ingredient.IngredientDetail(), func(i, j int) bool {
		return ingredient.IngredientDetail()[i].ExpirationDate.Before(ingredient.IngredientDetail()[j].ExpirationDate)
	})

	var ingredientStockID string
	for _, detail := range ingredient.IngredientDetail() {
		if detail.ExpirationDate.After(time.Now()) {
			ingredientStockID = detail.IngredientStockID

			if quantity >= detail.IngredientQuantity {
				err = s.ingredientRepo.DeleteUnexpiredIngredient(c, ingredientStockID)
				if err != nil {
					return err
				}
				quantity -= detail.IngredientQuantity
			} else {
				err = s.ingredientRepo.UpdateUnexpiredIngredientQuantity(c, ingredientStockID, detail.IngredientQuantity-quantity)
				if err != nil {
					return err
				}
				break
			}
		}
	}

	return nil
}

func (s *IngredientService) EditIngredient(c *fiber.Ctx, ingredient *domain.EditIngredientRequest) error {
	ingredientID := ingredient.IngredientID
	stockLessThan, _ := strconv.Atoi(ingredient.StockLessThan)
	dayBeforeExpire := util.ExpirationDate(ingredient.DayBeforeExpire)

	editIngredientPayload := &domain.EditIngredientPayload{
		IngredientID:       ingredientID,
		IngredientEngName:  ingredient.IngredientEngName,
		IngredientThaiName: ingredient.IngredientThaiName,
		StockLessThan:      stockLessThan,
		DayBeforeExpire:    dayBeforeExpire,
	}

	err := s.ingredientRepo.EditIngredient(c, editIngredientPayload)
	if err != nil {
		return err
	}

	return nil
}

func (s *IngredientService) GetAddEditIngredientDetail(c *fiber.Ctx, ingredientID string) (*domain.GetAddEditIngredientDetail, error) {
	ingredient, err := s.ingredientRepo.GetAddEditIngredientDetail(c, ingredientID)
	if err != nil {
		return nil, err
	}

	dayBeforeExpire := util.DaysSince2000(ingredient.DayBeforeExpire)
	dayBeforeExpireStr := strconv.Itoa(dayBeforeExpire)

	detail := &domain.GetAddEditIngredientDetail{
		IngredientEngName:  ingredient.IngredientEngName,
		IngredientThaiName: ingredient.IngredientThaiName,
		Unit:               string(ingredient.Unit),
		StockLessThan:      strconv.Itoa(int(ingredient.IngredientLessThan)),
		DayBeforeExpire:    dayBeforeExpireStr,
	}

	return detail, nil
}

func (s *IngredientService) EditIngredientStock(c *fiber.Ctx, ingredientStock *domain.EditIngredientStockRequest) error {
	ingredientStockID := ingredientStock.IngredientStockID
	price, _ := strconv.ParseFloat(ingredientStock.Price, 64)
	quantity, _ := strconv.ParseFloat(ingredientStock.Quantity, 64)
	expirationDate, _ := time.Parse("02/01/2006", ingredientStock.ExpirationDate)
	note := ingredientStock.Note

	editIngredientStockPayload := &domain.EditIngredientStockPayload{
		IngredientStockID: ingredientStockID,
		Price:             price,
		Quantity:          quantity,
		ExpirationDate:    expirationDate,
		Brand:             ingredientStock.Brand,
		Supplier:          ingredientStock.Supplier,
	}

	err := s.ingredientRepo.EditIngredientStock(c, editIngredientStockPayload)
	if err != nil {
		return err
	}

	if note != "" {
		ingredientNoteID := uuid.NewString()
		noteCreatedAt := time.Now()
		ingredientNote := &domain.AddIngredientNotePayload{
			IngredientNoteID:  ingredientNoteID,
			IngredientStockID: ingredientStockID,
			Note:              note,
			NoteCreatedAt:     noteCreatedAt,
		}

		err = s.ingredientRepo.AddIngredientNote(c, ingredientNote)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *IngredientService) GetEditIngredientStockDetail(c *fiber.Ctx, ingredientStockID string) (*domain.GetEditIngredientStockDetail, error) {
	ingredientStock, err := s.ingredientRepo.GetEditIngredientStockDetail(c, ingredientStockID)
	if err != nil {
		return nil, err
	}

	detail := &domain.GetEditIngredientStockDetail{
		IngredientStockID: ingredientStock.IngredientStockID,
		Brand:             ingredientStock.IngredientBrand,
		Quantity:          strconv.FormatFloat(ingredientStock.IngredientQuantity, 'f', -1, 64),
		Price:             strconv.FormatFloat(ingredientStock.Price, 'f', -1, 64),
		Supplier:          ingredientStock.IngredientSupplier,
		ExpirationDate:    ingredientStock.ExpirationDate.Format("02/01/2006"),
	}

	return detail, nil
}
