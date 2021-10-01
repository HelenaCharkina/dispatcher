package service

import "dispatcher/pkg/repository"

type Authorization interface {

}

type Agent interface {

}

type Service struct {
	Authorization
	Agent
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}