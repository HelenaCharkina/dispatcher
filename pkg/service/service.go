package service

import (
	"dispatcher/pkg/repository"
	"dispatcher/types"
)

type Authorization interface {
	GenerateToken(userId string) (string, error)
	GenerateRefreshToken() (string, error)
	SignIn(login string, password string) (*types.Response, error)
	ParseToken(token string) (string, error)
	SetRefreshToken(token string, userId string) error
	CheckRefreshToken(token string, userId string) error
	GenerateTokens(user *types.User) (*types.Response, error)
}

type Agent interface {
}


type Service struct {
	Authorization
	Agent
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization, repo.Token),
	}
}
