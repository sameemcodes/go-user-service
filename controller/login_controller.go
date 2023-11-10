// LoginController.go
package controller

import (
	"fmt"
	"net/http"

	"go-tools/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController(loginService *service.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
}

func (c *LoginController) GoogleLogin(g *gin.Context) {
	state := "random"
	url := c.loginService.LoginWithGoogle(state)
	g.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *LoginController) GoogleSignup(ctx *gin.Context) {
	// Get the ID token from the request body
	var request struct {
		IDToken string `json:"id_token"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Verify the ID token
	token, err := c.loginService.VerifyIDTokenUsingGoogle(request.IDToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the user info from the ID token
	userInfo, err := c.loginService.GetUserInfoUsingGoogle(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userInfo)

	// Create a new user with the obtained user info
	// Continue with your signup logic...
}
