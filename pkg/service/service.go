package service

import "github.com/ilgiz-ayupov/auth-app"

type Authorization interface {
	CreateUser(user auth.User) (int, error)
}

type Service struct {
	Authorization
}

func InitService() *Service {
	return &Service{
		Authorization: InitAuthService(),
	}
}
