package repository

import (
	"context"
	"go-tools/models"

	"go-tools/constants"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (uRepo *userRepository) GetUserByUserId(ctx context.Context, UserId string) (_ *models.User, err error) {
	var user models.User
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Where(constants.WhereUserId, UserId).Take(&user)
	return &user, getUser.Error
}

func (uRepo *userRepository) CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error) {
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Create(&user)
	return user, getUser.Error
}

func (uRepo *userRepository) GetAllUsers(ctx context.Context) (_ []models.User, err error) {
	var users []models.User
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Find(&users)
	return users, getUser.Error
}
func (uRepo *userRepository) Save(ctx context.Context, user *models.User) (_ *models.User, err error) {
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Where(constants.WhereUserId, user.UserId).Updates(user)
	return user, getUser.Error
}
func (uRepo *userRepository) DeleteUserById(ctx context.Context, userId string) (err error) {
	var user models.User
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Where(constants.WhereUserId, userId).Delete(&user)
	return getUser.Error
}
