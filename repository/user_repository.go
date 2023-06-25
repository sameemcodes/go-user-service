package repository

import (
	"context"
	"go-tools/models"
)

type UserRepository interface {
	GetUserByUserId(ctx context.Context, UserId string) (_ *models.User, err error)
	GetAllUsers(ctx context.Context) (_ []models.User, err error)
	CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error)
	Save(ctx context.Context, user *models.User) (_ *models.User, err error)
	DeleteUserById(ctx context.Context, userId string) (err error)
}
