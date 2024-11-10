package service

import (
	"github.com/BakingUp/BakingUp-Backend/internal/core/domain"
	"github.com/BakingUp/BakingUp-Backend/internal/core/port"
	"github.com/BakingUp/BakingUp-Backend/internal/core/util"
	"github.com/BakingUp/BakingUp-Backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAllUsers() ([]string, error) {
	usersDB, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var response []string

	for _, item := range usersDB {
		response = append(response, item.UserID)
	}

	return response, nil
}

func (s *UserService) GetUserInfo(c *fiber.Ctx, userID string) (*domain.UserInfo, error) {
	user, err := s.userRepo.GetUser(c, userID)
	if err != nil {
		return nil, err
	}

	orders, err := s.userRepo.GetUserProductionQueue(c, userID)
	if err != nil {
		return nil, err
	}

	language, err := s.GetUserLanguage(c, userID)
	if err != nil {
		return nil, err
	}

	var queue []domain.ProductionQueueItem
	for _, order := range orders {
		for _, orderProduct := range order.ProductionQueue() {
			recipe := orderProduct.Recipe()
			var recipeImg string
			if recipe != nil {
				for _, recipeImageItem := range recipe.RecipeImages() {
					if recipeImageItem.RecipeID == recipe.RecipeID {
						recipeImg = recipeImageItem.RecipeURL
					}
				}

				queue = append(queue, domain.ProductionQueueItem{
					OrderIndex: order.OrderIndex,
					Name:       util.GetRecipeName(recipe, language),
					Quantity:   orderProduct.ProductionQuantity,
					PickUpDate: order.PickUpDateTime.Format("02/01/2006"),
					ImgURL:     recipeImg,
				})
			}
		}
	}

	userInfo := &domain.UserInfo{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Tel:             user.Telephone,
		StoreName:       user.StoreName,
		ProductionQueue: queue,
	}

	return userInfo, nil
}

func (s *UserService) RegisterUser(user *domain.ManageUserRequest) (*domain.UserResponse, error) {
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return &domain.UserResponse{
			Status:  400,
			Message: "Cannot create a new user.",
		}, err
	}

	return &domain.UserResponse{
		Status:  201,
		Message: "Successfully create a new user.",
	}, nil
}

func (s *UserService) AddDeviceToken(req *domain.DeviceTokenRequest) (*domain.UserResponse, error) {
	err := s.userRepo.AddDeviceToken(req)
	if err != nil {
		return &domain.UserResponse{
			Status:  400,
			Message: "Cannot add a device token.",
		}, err
	}

	return &domain.UserResponse{
		Status:  201,
		Message: "Successfully added a device token.",
	}, nil
}

func (s *UserService) DeleteDeviceToken(req *domain.DeviceTokenRequest) (*domain.UserResponse, error) {
	err := s.userRepo.DeleteDeviceToken(req)
	if err != nil {
		return &domain.UserResponse{
			Status:  400,
			Message: "Cannot delete a device token.",
		}, err
	}

	return &domain.UserResponse{
		Status:  201,
		Message: "Successfully deleted a device token.",
	}, nil
}

func (s *UserService) DeleteAllExceptDeviceToken(req *domain.DeviceTokenRequest) (*domain.UserResponse, error) {
	err := s.userRepo.DeleteAllExceptDeviceToken(req)
	if err != nil {
		return &domain.UserResponse{
			Status:  400,
			Message: "Cannot delete all except this device token.",
		}, err
	}

	return &domain.UserResponse{
		Status:  201,
		Message: "Successfully delete all except this device token.",
	}, nil
}

func (s *UserService) GetUserLanguage(c *fiber.Ctx, userID string) (*db.Language, error) {
	user, err := s.userRepo.GetUser(c, userID)

	if err != nil {
		return nil, err
	}

	return &user.Language, nil
}

func (s *UserService) GetUserExpirationDate(c *fiber.Ctx, userID string) (*domain.ExpirationDate, error) {
	user, err := s.userRepo.GetUser(c, userID)

	if err != nil {
		return nil, err
	}

	expirationDate := &domain.ExpirationDate{
		YellowExpirationDate: user.YellowExpirationDate,
		RedExpirationDate:    user.RedExpirationDate,
		BlackExpirationDate:  user.BlackExpirationDate,
	}

	return expirationDate, nil
}

func (s *UserService) EditUserInfo(c *fiber.Ctx, editUserRequest *domain.ManageUserRequest) error {
	err := s.userRepo.EditUserInfo(c, editUserRequest)
	if err != nil {
		return err
	}

	return nil
}
