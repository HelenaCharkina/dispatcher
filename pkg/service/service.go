package service

import "dispatcher/pkg/repository"

type Authorization interface {
	GenerateToken(login string, password string) (string, error)
	ParseToken (token string) (string, error)
}

type Agent interface {

}

type Service struct {
	Authorization
	Agent
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}