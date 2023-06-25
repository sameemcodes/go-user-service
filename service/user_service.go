package service

import (
	"context"
	"go-tools/models"
)

// user service interface
type UserService interface {
	GetUserByUserId(ctx context.Context, UserId string) (_ *models.User, err error)
	GetAllUsers(ctx context.Context) (_ []models.User, err error)
	CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error)
	UpdateUser(ctx context.Context, user models.User) (_ *models.User, err error)
	DeleteUser(ctx context.Context, userId string) (err error)
}
