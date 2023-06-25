package controller

import (
	"go-tools/models"
	"go-tools/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserByUserId(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	CreateNewUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUserById(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// GetUserByUserId godoc
// @Summary Get a user by userId
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Param userId path string true "userId"
// @Router /user/v1/user_id/{userId} [get]
func (c *userController) GetUserByUserId(ctx *gin.Context) {

	userId := ctx.Params.ByName("userId")
	userDto, err := c.userService.GetUserByUserId(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetUserByUserId": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Router /user/v1/fetch_all [get]
func (c *userController) GetAllUsers(ctx *gin.Context) {
	userDto, err := c.userService.GetAllUsers(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetAllUsers": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// CreateNewUser godoc
// @Summary Create a new user
// @Tags User-Controller
// @Accept */*
// @Param user body models.User true "User"
// @Param userId path string true "userId"
// @Success 200
// @Failure 404
// @Failure 500
// @Produce json
// @Router /user/v1/create_new [post]
func (c *userController) CreateNewUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	userDto, err := c.userService.CreateNewUser(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during CreateNewUser ": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// UpdateUser godoc
// @Summary Update a user
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Router /user/v1/update [put]
func (c *userController) UpdateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)

	userDto, err := c.userService.UpdateUser(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during  UpdateUser": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// DeleteUserById godoc
// @Summary Delete a user by userId
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Success 200
// @Router /user/v1/deletebyUserId/{userId} [delete]
func (c *userController) DeleteUserById(ctx *gin.Context) {
	userId := ctx.Params.ByName("userId")
	err := c.userService.DeleteUser(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during DeleteUserById": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
