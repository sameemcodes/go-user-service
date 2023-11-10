// OAuthVendorService.go
package service

import (
	"context"
	"encoding/json"
	"errors"
	"go-tools/models"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type OAuthService interface {
	GetLoginURL(state string) string
	Exchange(code string) (*oauth2.Token, error)
	VerifyIDToken(idToken string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*models.UserInfo, error)
}

type GoogleService struct {
	conf *oauth2.Config
}

func NewGoogleService(clientID, clientSecret, redirectURL string) OAuthService {
	return &GoogleService{
		conf: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (s *GoogleService) GetLoginURL(state string) string {
	return s.conf.AuthCodeURL(state)
}

func (s *GoogleService) Exchange(code string) (*oauth2.Token, error) {
	return s.conf.Exchange(context.Background(), code)
}

func (s *GoogleService) VerifyIDToken(code string) (*oauth2.Token, error) {
	token, err := s.conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *GoogleService) GetUserInfo(token *oauth2.Token) (*models.UserInfo, error) {
	client := s.conf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo models.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

type FacebookService struct {
	config *oauth2.Config
}

func NewFacebookService(clientID, clientSecret, redirectURL string) OAuthService {
	return &FacebookService{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"public_profile", "email"},
			Endpoint:     facebook.Endpoint,
		},
	}
}

func (s *FacebookService) GetLoginURL(state string) string {
	return s.config.AuthCodeURL(state)
}

func (s *FacebookService) Exchange(code string) (*oauth2.Token, error) {
	return s.config.Exchange(context.Background(), code)
}

func (s *FacebookService) VerifyIDToken(idToken string) (*oauth2.Token, error) {
	return nil, errors.New("not implemented")
}

func (s *FacebookService) GetUserInfo(token *oauth2.Token) (*models.UserInfo, error) {
	return nil, errors.New("not implemented")
}

type GithubService struct {
	config *oauth2.Config
}

func NewGithubService(clientID, clientSecret, redirectURL string) OAuthService {
	return &GithubService{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"public_profile", "email"},
			Endpoint:     github.Endpoint,
		},
	}
}

func (s *GithubService) GetLoginURL(state string) string {
	return s.config.AuthCodeURL(state)
}

func (s *GithubService) Exchange(code string) (*oauth2.Token, error) {
	return s.config.Exchange(context.Background(), code)
}

func (s *GithubService) VerifyIDToken(idToken string) (*oauth2.Token, error) {
	return nil, errors.New("not implemented")
}
func (s *GithubService) GetUserInfo(token *oauth2.Token) (*models.UserInfo, error) {
	return nil, errors.New("not implemented")
}
