package service

import (
	"dispatcher/pkg/repository"
	"dispatcher/types"
)

type Authorization interface {
	SignIn(login string, password string) (*types.Response, error)
	RefreshToken(refreshToken string, userId string) (*types.Response, error)
	CheckToken(accessToken string) (string, error)
}

type Agent interface {
}


type Service struct {
	Authorization
	Agent
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.User, repo.Token),
	}
}
