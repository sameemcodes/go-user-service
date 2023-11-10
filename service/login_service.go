// LoginService.go
package service

import (
	"go-tools/models"

	"golang.org/x/oauth2"
)

type LoginService struct {
	googleService   OAuthService
	githubService   OAuthService
	facebookService OAuthService
}

func NewLoginService(googleService, githubService, facebookService OAuthService) *LoginService {
	return &LoginService{
		googleService:   googleService,
		githubService:   githubService,
		facebookService: facebookService,
	}
}

func (s *LoginService) LoginWithGoogle(state string) string {
	return s.googleService.GetLoginURL(state)
}

func (s *LoginService) SignupWithGoogle(code string) (*oauth2.Token, error) {
	return s.googleService.Exchange(code)
}

func (s *LoginService) VerifyIDTokenUsingGoogle(code string) (*oauth2.Token, error) {
	return s.googleService.VerifyIDToken(code)
}

func (s *LoginService) GetUserInfoUsingGoogle(token *oauth2.Token) (*models.UserInfo, error) {
	return s.googleService.GetUserInfo(token)
}

func (s *LoginService) LoginWithGithub(state string) string {
	return s.githubService.GetLoginURL(state)
}

func (s *LoginService) SignupWithGithub(code string) (*oauth2.Token, error) {
	return s.githubService.Exchange(code)
}

func (s *LoginService) VerifyIDTokenUsingGithub(code string) (*oauth2.Token, error) {
	return s.githubService.VerifyIDToken(code)
}

func (s *LoginService) GetUserInfoUsingGithub(token *oauth2.Token) (*models.UserInfo, error) {
	return s.githubService.GetUserInfo(token)
}

func (s *LoginService) LoginWithFacebook(state string) string {
	return s.facebookService.GetLoginURL(state)
}

func (s *LoginService) SignupWithFacebook(code string) (*oauth2.Token, error) {
	return s.facebookService.Exchange(code)
}

func (s *LoginService) VerifyIDTokenUsingFacebook(code string) (*oauth2.Token, error) {
	return s.facebookService.VerifyIDToken(code)
}

func (s *LoginService) GetUserInfoUsingFacebook(token *oauth2.Token) (*models.UserInfo, error) {
	return s.facebookService.GetUserInfo(token)
}
