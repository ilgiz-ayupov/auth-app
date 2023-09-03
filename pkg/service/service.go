package service

import (
	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (int64, error)
	GenerateJWTToken(login string, password string) (string, error)
	ParseJWTToken(token string) (auth.UserTokenClaims, error)
	AuthorizationToken(claims auth.UserTokenClaims) error
}

type Profile interface {
	GetUser(name string) (auth.UserProfile, error)
}

type Service struct {
	Authorization
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Profile:       NewProfileService(repos),
	}
}
