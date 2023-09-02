package service

import (
	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (int64, error)
	GenerateJWTToken(login string, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: InitAuthService(repos),
	}
}
