package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/repository"
)

const (
	salt       = "qfcv123leivn94"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type AuthService struct {
	repo *repository.Repository
}

func InitAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user auth.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateJWTToken(login string, password string) (string, error) {
	id, err := s.repo.AuthentificationUser(login, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
