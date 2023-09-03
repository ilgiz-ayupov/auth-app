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

func (s *ProfileService) AddPhoneNumber(phoneNum auth.PhoneNumber) (int64, error) {
	return s.repo.AddPhoneNumber(phoneNum)
}

func (s *ProfileService) SearchPhoneNumber(phone string) (auth.PhoneNumber, error) {
	return s.repo.SearchPhoneNumber(phone)
}

func (s *ProfileService) UpdatePhoneNumber(updating auth.UpdatingPhoneNumber) error {
	return s.repo.UpdatePhoneNumber(updating)
}

func (s *ProfileService) DeletePhoneNumber(id int) error {
	return s.repo.DeletePhoneNumber(id)
}
