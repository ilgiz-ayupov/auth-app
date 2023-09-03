package service

import (
	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/repository"
)

type ProfileService struct {
	repo *repository.Repository
}

func NewProfileService(repo *repository.Repository) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
}

func (s *ProfileService) GetUser(name string) (auth.UserProfile, error) {
	return s.repo.GetUser(name)
}

func (s *ProfileService) AddPhoneNumber(phoneNumber auth.PhoneNumber) (int64, error) {
	return s.repo.AddPhoneNumber(phoneNumber)
}
