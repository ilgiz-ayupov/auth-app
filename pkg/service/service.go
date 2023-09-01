package service

import (
	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (string, error)
}

type Service struct {
	Authorization
}

func InitService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: InitAuthService(repos),
	}
}
