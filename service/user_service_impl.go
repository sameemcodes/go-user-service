package service

import (
	"context"
	"fmt"
	"go-tools/models"
	"go-tools/repository"

	"github.com/mitchellh/mapstructure"
)

// user service implementation
type userService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) GetUserByUserId(ctx context.Context, userId string) (_ *models.User, err error) {
	entity, errorDb := service.userRepository.GetUserByUserId(ctx, userId)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *userService) GetAllUsers(ctx context.Context) (_ []models.User, err error) {
	entity, errorDb := service.userRepository.GetAllUsers(ctx)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *userService) CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error) {
	entity, errorDb := service.userRepository.CreateNewUser(ctx, user)
	if errorDb != nil {
		return user, errorDb
	}
	return entity, nil
}
func (service *userService) UpdateUser(ctx context.Context, user models.User) (_ *models.User, err error) {
	var userId = user.UserId
	newUser, err := service.GetUserByUserId(ctx, userId)
	fmt.Print("newUser  ", newUser, "err ", err)
	// Bind the JSON request body to the user object
	fmt.Print("user ", user)
	err2 := mapstructure.Decode(user, &newUser)
	fmt.Print("newUser", newUser, "user ", user)
	if err2 != nil {
		fmt.Print("err2  ", err2)
		return &user, err2
	}
	entity, errorDb := service.userRepository.Save(ctx, newUser)
	fmt.Print("entity  ", entity, "errorDb ", errorDb)
	if errorDb != nil {
		return &user, errorDb
	}
	return entity, nil
}
func (service *userService) DeleteUser(ctx context.Context, userId string) (err error) {
	//Delete User by Id
	errorDb := service.userRepository.DeleteUserById(ctx, userId)
	if errorDb != nil {
		return errorDb
	}
	return nil
}
